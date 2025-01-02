package handler

import (
	"api/models"
	"api/service"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// validate data
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
	var user UserRequest
	id_user := c.Locals("user_id").(uint)
	id,_ := strconv.Atoi(c.Params("id"))
	if uint(id) != id_user {
		return c.Status(403).JSON(fiber.Map{"message": "Forbidden"})
	}

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	userMap := map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
		"age":      user.Age,
		"phone":    user.Phone,
		"district": user.District,
		"city":     user.City,
		"state":    user.State,
		"country":  user.Country,
	}
	users:= service.UpdateUser(user, uint(id))
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
