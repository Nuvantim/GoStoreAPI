package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type (
	Permission = service.Permission //declare type models permission
)

/*
HANDLER Get Permission
*/
func GetPermission(c *fiber.Ctx) error {
	permission := service.GetPermission()
	if permission == nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Permission is empty",
		})
	}
	return c.Status(200).JSON(permission)
}

/*
HANDLER Find Permission
*/
func FindPermission(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	permission := service.FindPermission(uint64(id))
	if permission.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Permission Not Found",
		})
	}
	return c.Status(200).JSON(permission)
}

/*
HANDLER Create Permission
*/
func CreatePermission(c *fiber.Ctx) error {
	var permission Permission
	// bind body request
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid body request",
		})
	}
	// validate data
	if err := utils.Validator(permission); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	permissions := service.CreatePermission(permission)
	return c.Status(200).JSON(permissions)
}

/*
HANDLER Update Permission
*/
func UpdatePermission(c *fiber.Ctx) error {
	var permission Permission
	id, _ := c.ParamsInt("id")

	// check permission
	check_permission := service.FindPermission(uint64(id))
	if check_permission.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Permission Not Found",
		})
	}
	// bind body request
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid body reuqest",
		})
	}
	// validate data
	if err := utils.Validator(permission); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	permissions := service.UpdatePermission(uint64(id), permission)
	return c.Status(200).JSON(permissions)

}

/*
HANDLER Delete Permission
*/
func DeletePermission(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	permission, err := service.DeletePermission(uint64(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed delete permission",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": permission,
	})

}
