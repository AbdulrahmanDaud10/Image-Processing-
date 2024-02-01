package handlers

import (
	"fmt"
	"io"

	"github.com/AbdulrahmanDaud10/image-processing-golang-service/tasks"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("image")
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

	fileName := file.Filename
	resizeTasks, err := tasks.NewImageResizeTask(data, fileName)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create image resize tasks"})
	}

	client := tasks.GetClient()
	for _, task := range resizeTasks {
		if _, err := client.Enqueue(task); err != nil {
			fmt.Printf("Error enqueuing task: %v\n", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not enqueue image resize task"})
		}
	}

	return ctx.JSON(fiber.Map{"meaage": "Image uploaded and resizing task started"})
}
