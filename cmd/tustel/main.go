package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kobamkode/tustel"
	"github.com/kobamkode/tustel/db"
	"github.com/kobamkode/tustel/internal/env"
)

func main() {
	env.Load()

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	queries := db.New(pool)

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
		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		tustel.Process(file)

		if err := c.SaveFile(file, fmt.Sprintf("public/%s", file.Filename)); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		} else {
			if err := queries.CreateImageInfo(ctx, db.CreateImageInfoParams{
				Name: file.Filename,
				Path: fmt.Sprintf("public/%s", file.Filename),
			}); err != nil {
				return c.Status(fiber.StatusNotModified).SendString(err.Error())
			}

			return c.Status(fiber.StatusOK).SendString("file Uploaded")
		}

	})
	log.Fatal(app.Listen(":3000"))
}
