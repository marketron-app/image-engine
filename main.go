package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"marketron-image-engine/api/controllers"
	"marketron-image-engine/env"
	"os"
	"os/signal"
)

func init() {
	env.InitializeDotEnv()
}

func main() {
	app := fiber.New(fiber.Config{AppName: "Marketron Image Engine"})
	app.Use(recover.New())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	app.Get("/image", controllers.GetImage)

	if err := app.Listen(":3000"); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
}
