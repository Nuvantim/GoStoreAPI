package middleware

import (
	"github.com/gofiber/fiber/v3"
)

// Setup middleware function
func Setup() fiber.Handler {
	// methode middleware
	return func(c fiber.Ctx) error {
		return MethodMiddleware(c)
	}
    // JWT middleware
	return func(c fiber.Ctx) error {
		return AuthAndRefreshMiddleware(c)
	}
}
