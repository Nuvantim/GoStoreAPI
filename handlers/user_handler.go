package handler

import (
	"api/models"
	"api/service"
	"github.com/gofiber/fiber/v3"
)

// validate data
var user struct {
	name     string `json:"name" validate:"required"`
	email    string `json:"email" validate:"required"`
	password string `json:"password" validate:"required"`
	age      uint   `json:"age"`
	phone    uint   `json:"phone"`
	district string `json:"district"`
	city     string `json:"city"`
	state    string `json:"state"`
	country  string `json:"country"`
}

/*
Handler Get Profile
*/
func GetProfile(c fiber.Ctx) error {
	// Ambil User ID dari c.Locals
	userID := c.Locals("user_id").(uint)
	if userID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	// Query profil pengguna dari database
	user := service.FindUser(userID)

	return c.JSON(user)
}

/*
Handler GetUser
*/
func GetUser(c fiber.Ctx) error {
	user := service.GetUser()
	return c.Status(400).JSON(user)
}

/*
Handler Register User
*/
func RegisterUser(c fiber.Ctx) error {
	var users models.User
	if err := c.Bind().Body(&users); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	register := service.RegisterUser(users)
	return c.Status(200).JSON(register)
}

/*
Handler Update User
*/
func UpdateUser(c fiber.Ctx) error {
	id_user := c.Locals("user_id").(uint)
	id := strconv.Atoi(c.Params("id"))
	if uint(id) != id_user {
		return c.Status(403).JSON(fiber.Map{"message": "Forbidden"})
	}

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	users:= service.UpdateUser(user, id)
	return c.Status(200).JSON(users)
}

/*
Handler Delete User
*/
func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteUser(id); err != nil {
		return c.Status(500).SendString("Failed Delete User")
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "User Delete Succesfuly",
	})
}
