package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"

	"github.com/gofiber/fiber/v2"
)

type UserRole struct {
	RoleID []uint64 `json:"role_id" validate:"required"`
}

/*
Get Client
*/
func GetClient(c *fiber.Ctx) error {
	// start service
	client := service.GetClient()

	// return response data
	return c.Status(200).JSON(response.Pass("success get client", client))
}

/*
Find Client
*/
func FindClient(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	client := service.FindClient(uint64(id))

	// check client
	if client.ID == 0 {
		return c.Status(404).JSON(response.Error("failed find client", err.Error()))
	}

	// return response data
	return c.Status(200).JSON(response.Pass("success find client", client))
}

/*
Update Client
*/
func UpdateClient(c *fiber.Ctx) error {
	var req UserRole // declare variable strucf

	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// parser json
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(req); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}

	// start service
	client := service.UpdateClient(uint64(id), req.RoleID)

	// response data
	return c.Status(200).JSON(response.Pass("success update client", client))
}

/*
Remove Client
*/
func RemoveClient(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	client := service.FindClient(uint64(id))

	// check client
	if client.ID == 0 {
		return c.Status(404).JSON(response.Error("failed find client", "client not found"))
	}

	// start service
	if err := service.RemoveClient(uint64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed remove client", err.Error()))
	}

	// return response
	return c.Status(200).JSON(response.Pass("client removed", struct{}{}))
}
