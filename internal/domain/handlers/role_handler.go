package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	"github.com/gofiber/fiber/v2"
)

// struct request role
type role_permission struct {
	Name         string   `json:"name" validate:"required"`
	PermissionID []uint64 `json:"permission_id"`
}

type Role = service.Role //declare type role models

/*
HANDLER GET ROLE
*/
func GetRole(c *fiber.Ctx) error {
	role := service.GetRole()
	if role == nil {
		return c.Status(404).JSON(response.Error("failed get role", "role is empty"))
	}
	return c.Status(200).JSON(response.Pass("success get role", role))
}

/*
HANDLER FIND ROLE
*/
func FindRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id < 0 {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	role := service.FindRole(uint64(id))
	if role.ID == 0 {
		return c.Status(404).JSON(response.Error("failed find role", "role not found"))
	}
	return c.Status(200).JSON(response.Pass("success find role", role))
}

/*
HANDLER CREATE ROLE
*/
func CreateRole(c *fiber.Ctx) error {
	var req role_permission

	// bind body request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(req); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}
	role := service.CreateRole(req.Name, req.PermissionID)
	return c.Status(200).JSON(response.Pass("success create role", role))

}

/*
HANDLER UPDATE ROLE
*/
func UpdateRole(c *fiber.Ctx) error {
	var req role_permission
	id, err := c.ParamsInt("id")
	if err != nil || id < 0 {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	// bind body request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(req); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}
	role := service.UpdateRole(uint64(id), req.Name, req.PermissionID)
	return c.Status(200).JSON(response.Pass("success update role", role))
}

/*
HANDLER DELETE ROLE
*/
func DeleteRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id < 0 {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	if err := service.DeleteRole(uint64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete role", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("role deleted", struct{}{}))
}
