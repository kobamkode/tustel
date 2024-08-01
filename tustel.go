package tustel

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2/log"
)

func Process(file *multipart.FileHeader) {
	log.Infof("Processing %s", file.Filename)
}
