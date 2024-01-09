package handlers

import (
	"fmt"
	"io"

	"github.com/gofiber/fiber/v2"
)

func UploadImage(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Upload Failed"})
	}

	fileData, err := file.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open the file"})
	}
	defer fileData.Close()

	data, err := io.ReadAll(fileData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read the file"})
	}

	fmt.Println(data)

	// TODO: Resize the task

	return ctx.JSON(fiber.Map{"meaage": "Image upliaded and resizing task started"})
}
