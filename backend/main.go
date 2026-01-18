package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func main() {
	app := fiber.New()

	app.Use(healthcheck.New())
	log.Fatal(app.Listen(":3000"))
}
