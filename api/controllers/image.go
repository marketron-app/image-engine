package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	requestbody "marketron-image-engine/api/request-body"
	"marketron-image-engine/crawler"
	"marketron-image-engine/helpers"
	"os"
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

	err = os.WriteFile("tmp/"+fileName+"-temp.png", templateImage, 0644)
	err = os.WriteFile("tmp/"+fileName+".png", screenshotImage, 0644)

	return ctx.SendString("success")
}
