package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"

	"github.com/gofiber/fiber/v2"
)

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
	// declare variable struct
	var loginRequest = struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}{}

	// Parser Body JSON data
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(loginRequest); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}
	// start service
	accessToken, refreshToken, err := service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(401).JSON(response.Error("failed auth login", err.Error()))
	}

	// set cookies
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
	})

	// struct access token
	var loginResponse = struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: accessToken,
	}

	// response data
	return c.Status(200).JSON(response.Pass("success auth login login", loginResponse))
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
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
		MaxAge:   -1,
	})

	// Clear the refresh token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
		MaxAge:   -1, // Delete cookie
	})

	// response
	return c.Status(200).JSON(response.Pass("success logout", struct{}{}))
}

/*
Send OTP HANDLER
*/
func SendOTP(c *fiber.Ctx) error {
	// declare variable struct
	var otp = struct {
		Email string `json:"email" validate:"email,required"`
	}{}

	// parser body json
	if err := c.BodyParser(&otp); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}

	//validate data
	if err := validate.BodyStructs(otp); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}

	// start service
	send, err := service.SendOTP(otp.Email)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed send otp", err.Error()))
	}

	// response data
	return c.Status(200).JSON(response.Pass(send, struct{}{}))

}

func UpdatePassword(c *fiber.Ctx) error {
	// declare variable struct
	var password = struct {
		Otp      uint64 `json:"otp" validate:"required"`
		Password string `json:"password" validate:"required,min=8"`
	}{}

	// parser json body data
	if err := c.BodyParser(&password); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(password); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}

	// start service
	update, err := service.UpdatePassword(password.Otp, password.Password)
	if err != nil {
		return c.Status(500).JSON(response.Error("failed update password", err.Error()))
	}

	// response data
	return c.Status(200).JSON(response.Pass(update, struct{}{}))
}

