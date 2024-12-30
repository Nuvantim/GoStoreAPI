package handler

import (
	"api/service"
	"github.com/gofiber/fiber/v3"
	"api/utils"
)

// Fungsi untuk login atau refresh token secara otomatis
func Login(c fiber.Ctx) error {
	// Ambil refresh token dari header jika ada
	refreshToken := c.Cookies("refresh_token")
	if refreshToken != "" {
		// Jika ada refresh token, periksa dan buat access token baru
		newAccessToken, err := utils.RefreshAccessToken(refreshToken)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": err.Error()})
		}

		// Kirim access token baru
		return c.JSON(fiber.Map{
			"access_token": newAccessToken,
		})
	}

	// Jika tidak ada refresh token, lakukan login biasa
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON request ke struct
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Panggil service untuk login
	accessToken, refreshToken, err := service.Login(request.Email, request.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   true, // Set to true in production
		SameSite: "Strict",
		Path:     "/",
	})

	// Kirim access token dan refresh token
	return c.JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

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
