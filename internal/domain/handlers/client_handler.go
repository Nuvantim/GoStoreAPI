package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type UserRole struct {
	RoleID []uint64 `json:"role_id" validate:"required"`
}

/*
Get Client
*/
func GetClient(c *fiber.Ctx) error {
	client := service.GetClient()
	return c.Status(200).JSON(client)
}

/*
Find Client
*/
func FindClient(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	client := service.FindClient(uint64(id))
	// check client
	if client.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Client not found",
		})
	}
	return c.Status(200).JSON(client)

}

/*
Update Client
*/
func UpdateClient(c *fiber.Ctx) error {
	var req UserRole
	id, _ := c.ParamsInt("id")
	// bind body request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Invalid Body Request",
			"error":   err.Error(),
		})
	}

	// validate data
	if err := utils.Validator(req); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	client := service.UpdateClient(uint64(id), req.RoleID)
	return c.Status(200).JSON(client)
}

/*
Remove Client
*/
func RemoveClient(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	client := service.FindClient(uint64(id))
	// check client
	if client.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Client not found",
		})
	}

	if err := service.RemoveClient(uint64(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed delete client",
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Success remove client",
	})
}
