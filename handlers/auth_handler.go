package handler

import (
	"net/http"
	"toy-store-api/service"
	"github.com/gofiber/fiber/v2"
)

// Fungsi untuk login atau refresh token secara otomatis
func Login(c *fiber.Ctx) error {
	// Ambil refresh token dari header jika ada
	refreshToken := c.Get("Refresh-Token")
	if refreshToken != "" {
		// Jika ada refresh token, periksa dan buat access token baru
		newAccessToken, err := service.RefreshAccessToken(refreshToken)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
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
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Panggil service untuk login
	accessToken, refreshToken, err := service.Login(request.Email, request.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	// Kirim access token dan refresh token
	return c.JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
