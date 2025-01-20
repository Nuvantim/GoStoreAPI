package handler

import (
	"api/models"
	"api/service"
	"api/utils"
	"github.com/gofiber/fiber/v3"
)

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
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
	user := service.FindAccount(userID)

	return c.JSON(user)
}

/*
Handler GetUser
*/
// func GetUser(c fiber.Ctx) error {
// 	user := service.GetUser()
// 	return c.Status(400).JSON(user)
// }

/*
Handler Register User
*/
func RegisterAccount(c fiber.Ctx) error {
	var users models.User
	if err := c.Bind().Body(&users); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// validate data
	if err := utils.Validator(users); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//check email
	emails := service.CheckEmail(users.Email)
	if emails.ID != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Your email already exist",
		})
	}

	register := service.RegisterAccount(users)
	return c.Status(200).JSON(register)
}

/*
Handler Update User
*/
func UpdateAccount(c fiber.Ctx) error {
	var req UserRequest
	user_id := c.Locals("user_id").(uint)

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// validate data
	if err := utils.Validator(req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
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
	users := service.UpdateAccount(user, user_info, user_id)
	return c.Status(200).JSON(users)
}

/*
Handler Delete User
*/
func DeleteAccount(c fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint)
	if err := service.DeleteAccount(user_id); err != nil {
		return c.Status(500).SendString("Failed Delete User")
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "User Delete Succesfuly",
	})
}
