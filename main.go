package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	currentTime := time.Now()
	date := currentTime.Format("20060102")
	path := fmt.Sprintf("%s/%s-%s.%s", "logs", "log", date, "log")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	err := os.MkdirAll(filepath.Dir(path), 0770)
	if err != nil {
		logrus.Errorf("Failed to create log directory: %v", err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Errorf("Failed to open log file: %v", err)
		logrus.SetOutput(os.Stdout)
		return
	}
	defer file.Close() // Ensure file is closed even on errors
	logrus.SetOutput(file)

	// Create a new Fiber instance
	app := fiber.New()

	// Define a middleware for request logging
	app.Use(func(c *fiber.Ctx) error {
		logrus.WithFields(logrus.Fields{
			"app_name":    "myapp",
			"app_version": "1.0",
			"method":      c.Method(),
			"route":       c.Path(),
			"code":        c.Response().StatusCode(),
		}).Info("Request log")

		// Call the next handler in the chain
		return c.Next()
	})

	// Define a route for the root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, this is your API from kubernetes!")
	})

	// Start the Fiber app on port 9494
	err = app.Listen(":80")
	if err != nil {
		panic(err)
	}
}
