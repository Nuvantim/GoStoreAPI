package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// struct login
var login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// struct otp
var otp struct {
	Email string `json:"email" validate:"email,required"`
}

// struct update password
var password struct {
	Otp      uint64 `json:"otp" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

/*
Home Handler
*/
func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{})
}

/*
LOGIN HANDLER
*/
func Login(c *fiber.Ctx) error {
	// Bind data
	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Body Request"})
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
		"message":      "Login Success!",
		"access_token": accessToken,
	})
}

/*
LOGOUT HANDLER
*/
func Logout(c *fiber.Ctx) error {
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
Send OTP HANDLER
*/
func SendOTP(c *fiber.Ctx) error {
	// bind body
	if err := c.BodyParser(&otp); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid Body Request"})
	}
	send, err := service.SendOTP(otp.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{"message": send})

}

func UpdatePassword(c *fiber.Ctx) error {
	// bind body
	if err := c.BodyParser(&password); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid Body Request"})
	}

	// validate data
	if err := utils.Validator(password); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	update, err := service.UpdatePassword(password.Otp, password.Password)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"messsage": err.Error(),
		})
	}
	service.DeleteOTP(password.Otp)
	return c.Status(200).JSON(fiber.Map{
		"message": update,
	})
}
