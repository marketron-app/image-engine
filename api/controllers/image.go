package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	requestbody "marketron-image-engine/api/request-body"
	"marketron-image-engine/crawler"
	"marketron-image-engine/helpers"
	"marketron-image-engine/transformer"
	"marketron-image-engine/uploaders"
)

func GetImage(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	body := new(requestbody.GetImageBody)

	if err := ctx.QueryParser(body); err != nil {
		return err
	}

	fileName := uuid.New().String()

	err, templateImage := helpers.DownloadFromUrl(body.TemplateImage)
	if err != nil {
		ctx.SendString(err.Error())
		return ctx.SendStatus(500)
	}

	seleniumCrawler := crawler.Crawler{URL: body.URL, ViewportHeight: body.ViewportHeight, ViewportWidth: body.ViewportWidth}
	err, screenshotImage := seleniumCrawler.GetScreenshot()
	if err != nil {
		ctx.SendString(err.Error())
		return ctx.SendStatus(500)
	}

	trans := transformer.Transformer{WebsiteImage: screenshotImage, TemplateImage: templateImage, MappedCoordinates: body.Coordinates, FileName: fileName}
	err, finalImage := trans.Create()
	if err != nil {
		return err
	}

	uploaders.UploadToS3(fileName+".png", finalImage)

	return ctx.SendString("success")
}
