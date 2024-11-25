package handler

import (
	"github.com/gofiber/fiber/v3"
	"toy-store-api/service"
)

func Login(c fiber.Ctx) error {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind().Body(&credentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	token, err := service.Login(credentials.Email, credentials.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func Logout(c fiber.Ctx) error {
	// Logic for logout could be handled here, such as token invalidation
	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}

func Profile(c fiber.Ctx) error {
	// Get user information from context
	user := c.Locals("user")
	return c.JSON(fiber.Map{
		"user": user,
	})
}
