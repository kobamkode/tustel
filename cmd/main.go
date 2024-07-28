package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/kobamkode/tustel"
)

func main() {
	app := fiber.New()

	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} ${locals:requestid} ${status} - ${method} ${path} \"${resBody}\"\n",
	}))

	img := app.Group("/api/images")
	img.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	img.Post("/", func(c *fiber.Ctx) error {
		res, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		tustel.Process()

		return c.SendString(res.Filename)
	})
	log.Fatal(app.Listen(":3000"))
}
