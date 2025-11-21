package handler

import (
	"api/internal/domain/models"
	"api/internal/domain/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	"github.com/gofiber/fiber/v2"
)

type ( // declare type models User & UserInfo
	User     = models.User
	UserInfo = models.UserInfo
)

type userUpdate struct { //struct update Request
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"omitempty,min=8"`
	Age      uint64 `json:"age"`
	Phone    uint64 `json:"phone"`
	District string `json:"district"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
}

type userRegister struct {
	Otp      uint64 `json:"otp" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

/*
Handler Get Profile
*/
func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint64) // Get UserID from locals variable
	if userID == 0 {
		return c.Status(401).JSON(response.Error("failed get user id", "unauthorized"))
	}

	// Query user profile by id
	user, info := service.FindAccount(userID)

	data := struct {
		users     User     `json:"user"`
		user_info UserInfo `json:"user_info"`
	}{
		users:     user,
		user_info: info,
	}

	return c.Status(200).JSON(response.Pass("success get profile", data))
}

/*
Handler Register User
*/
func RegisterAccount(c *fiber.Ctx) error {
	var req userRegister

	// bind body data
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(req); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}

	// check email
	user_email := service.CheckEmail(req.Email)
	if user_email > 0 {
		return c.Status(409).JSON(response.Error("failed register account", "email already exist"))
	}

	// validate otp
	data, err := service.ValidateOTP(req.Otp)
	if err != nil {
		return c.Status(404).JSON(response.Error("failed validate otp", err.Error()))
	}

	// validate email
	if data.Email != req.Email {
		return c.Status(401).JSON(response.Error("failed register account", "register forbidden"))
	}

	// service register
	register, err := service.RegisterAccount(req.Name, req.Email, req.Password)
	if err != nil {
		return c.Status(400).JSON(response.Error("failed register account", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("success register account", register))

}

/*
Handler Update User
*/
func UpdateAccount(c *fiber.Ctx) error {
	var req userUpdate
	user_id := c.Locals("user_id").(uint64)
	user_email := c.Locals("email")

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(req); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}

	// check email
	if req.Email != user_email {
		user_email := service.CheckEmail(req.Email)
		if user_email > 0 {
			return c.Status(409).JSON(response.Error("failed update account", "email already exists"))
		}
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
		return c.Status(400).JSON(response.Error("failed update account", error.Error()))
	}
	// Make return interface
	data := struct {
		users     User     `json:"user"`
		user_info UserInfo `json:"user_info"`
	}{
		users:     users,
		user_info: userInfo,
	}

	return c.Status(200).JSON(response.Pass("success update account", data))
}

/*
Handler Delete User
*/
func DeleteAccount(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint64)
	if err := service.DeleteAccount(user_id); err != nil {
		return c.Status(500).JSON(response.Error("failed delete account", err.Error()))
	}

	return c.Status(201).JSON(response.Pass("account deleted", struct{}{}))
}
