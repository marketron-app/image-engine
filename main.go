package main

import "C"
import (
	"github.com/gofiber/fiber/v2"
	"marketron-image-engine/api/controllers"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "Marketron Image Engine"})

	app.Get("/image", controllers.GetImage)

	app.Listen(":3000")
}
