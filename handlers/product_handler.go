package handler

import (
	"e-commerce-api/models"
	"e-commerce-api/service"
	"github.com/gofiber/fiber/v3"
)

func GetProduct(c fiber.Ctx) error {
	product := service.GetAllProduct()
	return c.JSON(product)
}

func FindProduct(c fiber.Ctx) error {
	id := c.Params("id")
	product := service.GetProductById(id)
	return c.JSON(product)
}

func CreateProduct(c fiber.Ctx) error {
	var product models.Product
	if err := c.Bind().Body(product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	new_product:= service.CreateProduct(product)
	return c.JSON(new_product)
}


