package routes

import "github.com/gofiber/fiber/v2"

func RouteSetup(app *fiber.App) {
	app.Post("/image-process", handlers.UploadImage)
}