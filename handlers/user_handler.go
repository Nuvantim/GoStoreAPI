package handler

import (
	"e-commerce-api/service"
	"github.com/gofiber/fiber/v3"
)

func GetUser(c fiber.Ctx) error {
	user := service.GetUser()
	return c.Status(200).JSON(user)
}

func FindUser(c fiber.Ctx) error {
	id := Params("id")
	user := service.FindUser(id)
	return c.Status(200).JSON(user)
}

func RegisterUser(c fiber.Ctx) error {
	// validate data
	var user struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
		Address  string `json:"address" validate:"required"`
		Phone    uint   `json:"phone" validate:"required"`
	}

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	users := service.RegisterUser(user.Name, user.Email, user.Password, user.Address, user.Phone)
	return c.Status(200).JSON(user)
}
