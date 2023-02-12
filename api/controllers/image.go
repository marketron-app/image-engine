package controllers

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	requestbody "marketron-image-engine/api/request-body"
	"marketron-image-engine/crawler"
	"marketron-image-engine/helpers"
	"marketron-image-engine/transformer"
	"marketron-image-engine/uploaders"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	crawlerTimeMetricHeaderName     = "X-MARKETRON-METRIC-CRAWLER"
	transformerTimeMetricHeaderName = "X-MARKETRON-METRIC-TRANSFORMER"
	uploaderTimeMetricHeaderName    = "X-MARKETRON-METRIC-UPLOADER"
	defaultCrawlerTimeout           = 10
)

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

	contextLogger := log.WithFields(log.Fields{
		"url":           body.URL,
		"templateImage": body.TemplateImage,
	})

	fileName := body.FileName

	if fileName == "" {
		fileName = uuid.New().String()
	}

	err, templateImage := helpers.DownloadFromUrl(body.TemplateImage)
	if err != nil {
		contextLogger.Error("Error downloading template: " + err.Error())
		ctx.SendString(err.Error())
		return ctx.SendStatus(500)
	}

	start := time.Now()
	seleniumCrawler := crawler.Crawler{URL: body.URL, ViewportHeight: body.ViewportHeight, ViewportWidth: body.ViewportWidth, IsMobile: body.IsMobile}
	contextTimeoutSeconds, err := strconv.Atoi(os.Getenv("CRAWLER_TIMEOUT"))
	if err != nil || contextTimeoutSeconds == 0 {
		contextTimeoutSeconds = defaultCrawlerTimeout
	}

	timeoutCtx, cancel := context.WithTimeout(ctx.Context(), time.Duration(contextTimeoutSeconds)*time.Second)
	defer cancel()

	err, screenshotImage := seleniumCrawler.GetScreenshot(timeoutCtx)
	if err != nil {
		contextLogger.Error("Error getting screenshot: " + err.Error())
		ctx.SendString(err.Error())
		return ctx.SendStatus(500)
	}
	crawlerTime := time.Since(start).Milliseconds()
	addMetricHeader(ctx, crawlerTimeMetricHeaderName, fmt.Sprintf("%d", crawlerTime))

	start = time.Now()
	trans := transformer.Transformer{WebsiteImage: screenshotImage, TemplateImage: templateImage, MappedCoordinates: body.Coordinates, FileName: fileName}
	err, finalImage := trans.Create()
	if err != nil {
		contextLogger.Error("Error transforming image: " + err.Error())
		return err
	}
	transformerTime := time.Since(start).Milliseconds()
	addMetricHeader(ctx, transformerTimeMetricHeaderName, fmt.Sprintf("%d", transformerTime))

	start = time.Now()
	err = uploaders.UploadToS3(ctx.Context(), fileName+".png", finalImage)
	if err != nil {
		log.Error("Error uploading to S3: " + err.Error())
	}
	uploaderTime := time.Since(start).Milliseconds()
	addMetricHeader(ctx, uploaderTimeMetricHeaderName, fmt.Sprintf("%d", uploaderTime))

	return ctx.JSON(fiber.Map{
		"filename": fileName + ".png",
	})
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(imageBody requestbody.GetImageBody) []*ErrorResponse {
	var errors []*ErrorResponse
	var validate = validator.New()
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

func addMetricHeader(ctx *fiber.Ctx, headerName, headerValue string) {
	var headersEnabled = os.Getenv("METRIC_HEADERS_ENABLED")
	if strings.ToLower(headersEnabled) == "true" {
		ctx.Append(headerName, headerValue)
	}
}
