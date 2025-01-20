package handler

import (
	"api/models"
	"api/service"
	"api/utils"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

/*
HANDLER GET CATEGORY
*/
func GetCategory(c fiber.Ctx) error {
	category := service.GetAllCategory()
	return c.Status(200).JSON(category)
}

/*
HANDLER FIND CATEGORY
*/
func FindCategory(c fiber.Ctx) error {
	id := c.Params("id")
	category := service.GetCategoryById(id)
	return c.Status(200).JSON(category)
}

/*
HANDLER CREATE CATEGORY
*/
func CreateCategory(c fiber.Ctx) error {
	var category models.Category
	if err := c.Bind().Body(&category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validate data
	if err := utils.Validator(category); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	categories := service.CreateCategory(category)
	return c.Status(200).JSON(categories)
}

/*
HANDLER UPDATE CATEGORY
*/
func UpdateCategory(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var category models.Category
	if err := c.Bind().Body(&category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validate data
	if err := utils.Validator(category); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	category_update := service.UpdateCategory(uint(id), category)
	return c.Status(200).JSON(category_update)
}

/*
HANDLER DELETE CATEGORY
*/
func DeleteCategory(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := service.DeleteCategory(uint(id)); err != nil {
		return c.Status(500).SendString("Failed Delete Category")
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
