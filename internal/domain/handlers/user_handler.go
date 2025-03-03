package handler

import (
	"api/internal/domain/models"
	"api/internal/domain/service"
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ( // declare type models User & UserTemps
	User     = models.User
	UserTemp = models.UserTemp
	UserInfo = models.UserInfo
)

type UserRequest struct { //struct update Request
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"omitempty,min=8"`
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
func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint) // Get UserID from locals variable
	if userID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	// Query user profile by id
	user, info := service.FindAccount(userID)

	return c.Status(200).JSON(fiber.Map{
		"user":      user,
		"user_info": info,
	})
}

/*
Handler Register User
*/
func RegisterAccount(c *fiber.Ctx) error {
	var users UserTemp
	if err := c.BodyParser(&users); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// validate data
	if err := utils.Validator(users); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//check email
	email_user, email_usertemp := service.CheckEmail(users.Email)
	if email_user > 0 || email_usertemp > 0 {
		return c.Status(409).JSON(fiber.Map{
			"message": "Email already exist",
		})
	}

	register := service.RegisterAccount(users)
	return c.Status(200).JSON(fiber.Map{
		"message": register,
	})
}

/*
Handler Update User
*/
func UpdateAccount(c *fiber.Ctx) error {
	var req UserRequest
	user_id := c.Locals("user_id").(uint)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// validate data
	if err := utils.Validator(req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// check email
	user_email, _ := service.CheckEmail(req.Email)
	if user_email > 0 {
		return c.Status(409).JSON(fiber.Map{
			"message": "Email is already exist in another user",
		})
	}

	//parsing user model
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	//parsing user_info model
	user_info := UserInfo{
		Age:      req.Age,
		Phone:    req.Phone,
		District: req.District,
		City:     req.City,
		State:    req.State,
		Country:  req.Country,
	}
	users, userInfo, error := service.UpdateAccount(user, user_info, user_id)
	if error != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": error,
		})
	}
	// Make return interface
	return c.Status(200).JSON(fiber.Map{
		"user":      users,
		"user_info": userInfo,
	})
}

/*
Handler Delete User
*/
func DeleteAccount(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint)
	if err := service.DeleteAccount(user_id); err != nil {
		return c.Status(500).SendString("Failed Delete User")
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "User Delete Succesfuly",
	})
}
