package handler

import (
	"api/models"
	"api/service"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      uint   `json:"age"`
	Phone    uint   `json:"phone"`
	District string `json:"district"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
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
	var req UserRequest
	id_user := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))
	if uint(id) != id_user {
		return c.Status(403).JSON(fiber.Map{"message": "Forbidden"})
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	//parsing user model
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	//parsing user_info model
	user_info := models.UserInfo{
		Age:      req.Age,
		Phone:    req.Phone,
		District: req.District,
		City:     req.City,
		State:    req.State,
		Country:  req.Country,
	}
	users := service.UpdateUser(user, user_info, uint(id))
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
