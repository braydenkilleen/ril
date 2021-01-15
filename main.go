package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// app := fiber.New()
	// app.Use(cors.New())
	// setupRoutes(app)
	app := Setup()
	log.Fatal(app.Listen(":3000"))
}

// Setup sets up an App instance
func Setup() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	return app
}

func setupRoutes(app *fiber.App) {
	app.Get("/", index)
}

func index(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
