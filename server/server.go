package main

import (
	"log"

	"github.com/AbdulrahmanDaud10/image-processing-golang-service/routes"
	"github.com/gofiber/fiber/v2"
)

const redisAddress = "127.0.0.1:6379"

func main() {
	app := fiber.New()
	routes.RouteSetup(app)
	log.Fatal(app.Listen(":3000"))
}
