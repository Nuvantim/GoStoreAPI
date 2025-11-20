package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"

	"github.com/gofiber/fiber/v2"
)

type Category = service.Category //declare type model Category
/*
HANDLER GET CATEGORY
*/
func GetCategory(c *fiber.Ctx) error {
	// start service
	category := service.GetAllCategory()

	// return response data
	return c.Status(200).JSON(response.Pass("success get category", category))
}

/*
HANDLER FIND CATEGORY
*/
func FindCategory(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	category := service.FindCategory(uint64(id))

	// return response data
	return c.Status(200).JSON(response.Pass("success find category", category))
}

/*
HANDLER CREATE CATEGORY
*/
func CreateCategory(c *fiber.Ctx) error {
	var category Category // declare variabel Category

	// parser json
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(category); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}

	// start service
	categories := service.CreateCategory(category)

	// return response data
	return c.Status(200).JSON(response.Pass("success create category", categories))
}

/*
HANDLER UPDATE CATEGORY
*/
func UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // get params ID
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	var category Category // declare variabel Category

	// parser json
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}
	// validate data
	if err := validate.BodyStructs(category); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}
	// start service
	category_update := service.UpdateCategory(uint64(id), category)

	// return response data
	return c.Status(200).JSON(response.Pass("success update category", category_update))
}

/*
HANDLER DELETE CATEGORY
*/
func DeleteCategory(c *fiber.Ctx) error {
	// get id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// start service
	if err := service.DeleteCategory(uint64(id)); err != nil {
		return c.Status(500).JSON(response.Error("delete category", err.Error()))
	}

	// return response
	return c.Status(200).JSON(response.Pass("category deleted", struct{}{}))
}
