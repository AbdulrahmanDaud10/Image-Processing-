package routes

import (
	"github.com/AbdulrahmanDaud10/image-processing-golang-service/handlers"
	"github.com/gofiber/fiber/v2"
)

func RouteSetup(app *fiber.App) {
	app.Post("/image-process", handlers.UploadImage)
}
