package handler

import (
	"api/service"
	"api/utils"
	"github.com/gofiber/fiber/v3"
)

// struct login
var login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// struct otp
var otp struct {
	Code string `json:"code"`
}

/*
LOGIN HANDLER
*/
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

	// Bind data
	if err := c.Bind().Body(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// validate data
	if err := utils.Validator(login); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Panggil service untuk login
	accessToken, refreshToken, err := service.Login(login.Email, login.Password)
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
		"Message": "Login Success!",
		"access_token":  accessToken,
	})
}

/*
LOGOUT HANDLER
*/
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

/*
vERIFY OTP HANDLER
*/
func OtpVerify(c fiber.Ctx) error {
	// bind body
	if err := c.Bind().Body(&otp); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}
	verify := service.OtpVerify(otp.Code)

	return c.Status(200).JSON(fiber.Map{"message": verify})

}
