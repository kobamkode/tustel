package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kobamkode/tustel/internal/routes"
)

func Run(pool *pgxpool.Pool) {
	r := routes.NewHandle(pool)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} ${locals:requestid} ${latency} ${status} - ${method} ${path} \"${resBody}\"\n",
	}))

	img := app.Group("/api/images")
	img.Get("/", r.GetImage)
	img.Post("/", r.UploadImage)
	log.Fatal(app.Listen(":3000"))
}
