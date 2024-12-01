package middleware

import (
	"github.com/gofiber/fiber/v3"
)

// Setup is the middleware that sets up JWT authentication
func Setup() fiber.Handler {
	// Return a handler that uses JwtToken middleware
	return func(c fiber.Ctx) error {
		// Apply the JWT token middleware
		return AuthAndRefreshMiddleware(c)
	}
}
