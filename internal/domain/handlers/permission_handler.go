package handler

import (
	"api/internal/domain/service"
	"api/pkg/guard"
	"github.com/gofiber/fiber/v2"
    "api/pkg/utils/responses"
    "api/pkg/utils/validates"
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
		return c.Status(404).JSON(response.Error("failed get permission", "permission is empty"))
	}
	return c.Status(200).JSON(response.Pass("success get permission", permission))
}

/*
HANDLER Find Permission
*/
func FindPermission(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	permission := service.FindPermission(uint64(id))
	if permission.ID == 0 {
		return c.Status(404).JSON(response.Error("failed find permission", "permission not found"))
	}
	return c.Status(200).JSON(response.Pass("success find permission", permission))
}

/*
HANDLER Create Permission
*/
func CreatePermission(c *fiber.Ctx) error {
	var permission Permission
	// bind body request
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(permission); err != nil{
        return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
    }
    
	permissions := service.CreatePermission(permission)
	return c.Status(200).JSON(response.Pass("success create permission", permission))
}

/*
HANDLER Update Permission
*/
func UpdatePermission(c *fiber.Ctx) error {
	var permission Permission
	id, err := c.ParamsInt("id")
    if err != nil{
        return c.Status(400).JSON(response.Error("failed get id", err.Error()))
    }

	// check permission
	check_permission := service.FindPermission(uint64(id))
	if check_permission.ID == 0 {
		return c.Status(404).JSON(response.Error("failed find permission", "permission not found"))
	}
	// bind body request
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(permission); err != nil{
        return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
    }

	permissions := service.UpdatePermission(uint64(id), permission)
	return c.Status(200).JSON(response.Pass("success update permission", permissions))

}

/*
HANDLER Delete Permission
*/
func DeletePermission(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
    if err != nil{
        return c.Status(400).JSON("failed get id", err.Error())
    }
	permission, err := service.DeletePermission(uint64(id))
	if err != nil {
		return c.Status(500).JSON(response.Error("failed delete permission", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("permission deleted", struct{}{}))

}
