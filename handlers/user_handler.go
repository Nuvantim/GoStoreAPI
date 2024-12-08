package handler

import (
	"api/database"
	"api/models"
	"api/service"
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

func GetProfile(c fiber.Ctx) error {
	// Ambil User ID dari c.Locals
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Query profil pengguna dari database
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func GetUser(c fiber.Ctx) error {
	user := service.GetUser()
	return c.Status(400).JSON(user)
}

func RegisterUser(c fiber.Ctx) error {
	var users models.User
	if err := c.Bind().Body(&users); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	service.RegisterUser(users)
	return c.Status(200).JSON(users)
}

func FindUser(c fiber.Ctx) error {
	id := c.Params("id")
	user := service.FindUser(id)
	return c.Status(200).JSON(user)
}

func UpdateUser(c fiber.Ctx) error {
	var users models.User
	id_user := c.Locals("user_id").(string)
	id := c.Params("id")
	if id != id_user {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	if err := c.Bind().Body(&users); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	service.UpdateUser(id, users)
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
