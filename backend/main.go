package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if code == fiber.StatusNotFound {
				return c.Status(code).SendString(http.StatusText(code))
			}

			c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
			return c.Status(code).SendString(err.Error())
		},
	})

	app.Use(healthcheck.New())

	log.Fatal(app.Listen(":3000"))
}
