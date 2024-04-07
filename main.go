package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define a route for the root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, this is your API from kubernetes!")
	})

	// Start the Fiber app on port 9494
	err := app.Listen(":80")
	if err != nil {
		panic(err)
	}
}
