package handlers

import (
	"e-commerce-api/models"
	"e-commerce-api/service"
	"github.com/gofiber/fiber/v3"
)

func GetCategory(c Fiber.Ctx) error {
	category := service.GetAllCategory()
	return c.JSON(category)
}

func FindCategory(c Fiber.Ctx) error {
	id := c.Params("id")
	category := service.GetCategoryById(id)
	return c.JSON(category)
}

func CreateCategory(c Fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil{
		return c.Status(400).SendString(err.Error())
	}
	new_category:= service.CreateCategory(category)
	return c.JSON(new_category)
}

func UpdateCategory (c fiber.Ctx) error {
	id := c.Params("id")
	var category_data models.Category
	if err := c.BodyParser(&category_data); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	update_category := service.UpdateCategory(id, category_data)
	return update_category
}

func DeleteCategory (c fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteCategory(id); err != nil {
		return c.Status(500).SendString("Failed Delete Category")
	}
	return c.SendStatus(204)
}


