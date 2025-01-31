package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func MethodMiddleware(c fiber.Ctx) error {
	methods := c.Route().Method
	if c.Method() != methods {
		return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"error":   "HTTP method not allowed",
			"allowed": methods,
		})
	}
	return c.Next()
}
