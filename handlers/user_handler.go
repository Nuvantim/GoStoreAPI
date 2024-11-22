package handler

import (
	"e-commerce-api/service"
	"github.com/gofiber/fiber/v3"
)

// validate data
var user struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Phone    uint   `json:"phone" validate:"required"`
}

func GetUser(c fiber.Ctx) error {
	user := service.GetUser()
	return c.Status(200).JSON(user)
}

func FindUser(c fiber.Ctx) error {
	id := c.Params("id")
	user := service.FindUser(id)
	return c.Status(200).JSON(user)
}

func RegisterUser(c fiber.Ctx) error {
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	users := service.RegisterUser(user.Name, user.Email, user.Password, user.Address, user.Phone)
	return c.Status(200).JSON(users)
}

func UpdateUser(c fiber.Ctx) error {
	id := c.Params("id")
	if err := c.Bind().Body(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	users := service.UpdateUser(id, user.Name, user.Email, user.Password, user.Address, user.Phone)
	return c.Status(200).JSON(users)
}

func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteUser(id); err != nil {
		return c.Status(500).SendString("Failed Delete User")
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "User Delete Succesfuly",
	})
}
