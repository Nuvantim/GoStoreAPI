package handler

import (
	"api/service"
	"api/utils"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

var product service.Product //declare variabel Product

func GetProduct(c fiber.Ctx) error {
	products := service.GetAllProduct()
	return c.Status(200).JSON(products)
}

func FindProduct(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := service.FindProduct(uint(id))
	return c.Status(200).JSON(product)
}

func CreateProduct(c fiber.Ctx) error {
	if err := c.Bind().Body(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// validate data
	if err := utils.Validator(product); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// check category
	ctg := service.FindCategory(product.CategoryID)
	if ctg.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Message": "Category not found",
		})
	}

	service.CreateProduct(product)
	return c.Status(200).JSON(product)
}

func UpdateProduct(c fiber.Ctx) error {
	id,_ := strconv.Atoi(c.Params("id"))
	if err := c.Bind().Body(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validate data
	if err := utils.Validator(product); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	products := service.UpdateProduct(uint(id), product)
	return c.Status(200).JSON(products)
}

func DeleteProduct(c fiber.Ctx) error {
	id,_ := strconv.Atoi(c.Params("id"))
	if err := service.DeleteProduct(uint(id)); err != nil {
		return c.Status(500).SendString("Failed Delete Product")
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
