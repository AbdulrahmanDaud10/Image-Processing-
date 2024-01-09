package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

const redisAddress = "127.0.0.1:6379"

func main() {
	app := fiber.New()
	log.Fatal(app.Listen(":3000"))
}