package controllers

import "github.com/gofiber/fiber/v2"

func Health(ctx *fiber.Ctx) error {
	return ctx.SendStatus(200)
}
