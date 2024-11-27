package handler

import (
	"github.com/gofiber/fiber/v3"
	"toy-store-api/service"
)

// LoginHandler handles user login and sets cookies
func Login(c fiber.Ctx) error {
	// Struct for request body
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Authenticate user and generate access token (no refresh token)
	accessToken, err := service.Login(req.Email, req.Password) // Expecting only 2 return values
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Set the access token in a cookie
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HTTPOnly: true,
		Secure:   true, // Set to true in production
		SameSite: "Strict",
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"message": "Login successful",
	})
}

// LogoutHandler clears the access and refresh tokens from cookies
func Logout(c fiber.Ctx) error {
	// Clear the access token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   true, // Set to true in production
		SameSite: "Strict",
		Path:     "/",
		MaxAge:   -1, // Delete cookie
	})

	// Clear the refresh token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   true, // Set to true in production
		SameSite: "Strict",
		Path:     "/",
		MaxAge:   -1, // Delete cookie
	})

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}
