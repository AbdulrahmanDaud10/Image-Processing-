package main

import (
	"log"

	"github.com/AbdulrahmanDaud10/image-processing-golang-service/routes"
	"github.com/AbdulrahmanDaud10/image-processing-golang-service/tasks"
	"github.com/gofiber/fiber/v2"
)

const redisAddress = "127.0.0.1:6379"

func main() {
	app := fiber.New()
	routes.RouteSetup(app)
	tasks.InitRedis(redisAddress)
	defer tasks.Close()
	log.Fatal(app.Listen(":3000"))
}
