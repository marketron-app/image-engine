package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	requestbody "marketron-image-engine/api/request-body"
	"marketron-image-engine/crawler"
	"marketron-image-engine/helpers"
	"marketron-image-engine/transformer"
	"marketron-image-engine/uploaders"
	"time"
)

var validate = validator.New()

func GetImage(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	body := new(requestbody.GetImageBody)

	if err := ctx.QueryParser(body); err != nil {
		return err
	}

	errors := ValidateStruct(*body)
	if errors != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)

	}

	fileName := uuid.New().String()

	err, templateImage := helpers.DownloadFromUrl(body.TemplateImage)
	if err != nil {
		ctx.SendString(err.Error())
		return ctx.SendStatus(500)
	}

	start := time.Now()
	seleniumCrawler := crawler.Crawler{URL: body.URL, ViewportHeight: body.ViewportHeight, ViewportWidth: body.ViewportWidth}
	err, screenshotImage := seleniumCrawler.GetScreenshot()
	if err != nil {
		ctx.SendString(err.Error())
		return ctx.SendStatus(500)
	}
	fmt.Println("Crawler time: " + time.Since(start).String())

	start = time.Now()
	trans := transformer.Transformer{WebsiteImage: screenshotImage, TemplateImage: templateImage, MappedCoordinates: body.Coordinates, FileName: fileName}
	err, finalImage := trans.Create()
	if err != nil {
		return err
	}
	fmt.Println("Transformer time: " + time.Since(start).String())

	start = time.Now()
	uploaders.UploadToS3(fileName+".png", finalImage)
	fmt.Println("Uploader time: " + time.Since(start).String())

	return ctx.SendString("success")
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(imageBody requestbody.GetImageBody) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(imageBody)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
