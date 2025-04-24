package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type Category = service.Category //declare type model Category
/*
HANDLER GET CATEGORY
*/
func GetCategory(c *fiber.Ctx) error {
	category := service.GetAllCategory()
	return c.Status(200).JSON(category)
}

/*
HANDLER FIND CATEGORY
*/
func FindCategory(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id") // get params ID
	category := service.FindCategory(uint64(id))
	return c.Status(200).JSON(category)
}

/*
HANDLER CREATE CATEGORY
*/
func CreateCategory(c *fiber.Ctx) error {
	var category Category // declare variabel Category
	if err := c.BodyParser(&category); err != nil {
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
func UpdateCategory(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id") // get params ID
	var category Category      // declare variabel Category

	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validate data
	if err := utils.Validator(category); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	category_update := service.UpdateCategory(uint64(id), category)
	return c.Status(200).JSON(category_update)
}

/*
HANDLER DELETE CATEGORY
*/
func DeleteCategory(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := service.DeleteCategory(uint64(id)); err != nil {
		return c.Status(500).SendString("Failed Delete Category")
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
