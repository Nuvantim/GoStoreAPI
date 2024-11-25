package handler

import (
	"github.com/gofiber/fiber/v3"
	"toy-store-api/models"
	"toy-store-api/service"
)

func GetCategory(c fiber.Ctx) error {
	category := service.GetAllCategory()
	return c.Status(200).JSON(category)
}

func FindCategory(c fiber.Ctx) error {
	id := c.Params("id")
	category := service.GetCategoryById(id)
	return c.Status(200).JSON(category)
}

func CreateCategory(c fiber.Ctx) error {
	var category models.Category
	if err := c.Bind().Body(&category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	categories := service.CreateCategory(category)
	return c.Status(200).JSON(categories)
}

func UpdateCategory(c fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := c.Bind().Body(&category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	service.UpdateCategory(id, category)
	return c.Status(200).JSON(category)
}

func DeleteCategory(c fiber.Ctx) error {
	id := c.Params("id")
	if err := service.DeleteCategory(id); err != nil {
		return c.Status(500).SendString("Failed Delete Category")
	}
	return c.SendStatus(204)
}
