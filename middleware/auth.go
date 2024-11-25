package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func JWTMiddleware(c fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	// Remove 'Bearer ' prefix
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Parse and validate the JWT token
	claims, err := service.ParseToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Set user info in context
	c.Locals("user", claims)
	return c.Next()
}
