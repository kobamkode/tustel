package routes

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kobamkode/tustel"
	"github.com/kobamkode/tustel/db"
)

type Handle struct {
	pool *pgxpool.Pool
}

func NewHandle(pool *pgxpool.Pool) *Handle {
	return &Handle{
		pool: pool,
	}
}

func (h *Handle) UploadImage(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	tustel.Process(file)

	if err := c.SaveFile(file, fmt.Sprintf("public/%s", file.Filename)); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	} else {
		if err := db.New(h.pool).CreateImageInfo(ctx, db.CreateImageInfoParams{
			Name: file.Filename,
			Path: fmt.Sprintf("public/%s", file.Filename),
		}); err != nil {
			return c.Status(fiber.StatusNotModified).SendString(err.Error())
		}

		return c.Status(fiber.StatusOK).SendString("file Uploaded")
	}
}

func (h *Handle) GetImage(c *fiber.Ctx) error {
	return nil

}
