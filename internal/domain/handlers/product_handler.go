package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type Product = service.Product //declare type models Product

func GetProduct(c *fiber.Ctx) error {
	products := service.GetAllProduct()
	return c.Status(200).JSON(products)
}

func FindProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product := service.FindProduct(uint64(id))
	return c.Status(200).JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	var product Product //declare variabel Product
	if err := c.BodyParser(&product); err != nil {
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
			"message": "Category not found",
		})
	}

	service.CreateProduct(product)
	return c.Status(200).JSON(product)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id") // get Params ID
	var product Product        //declare variabel Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validate data
	if err := utils.Validator(product); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	products := service.UpdateProduct(uint64(id), product)
	return c.Status(200).JSON(products)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := service.DeleteProduct(uint64(id)); err != nil {
		return c.Status(500).SendString("Failed Delete Product")
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
