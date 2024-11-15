package handlers

import (
	"e-commerce-api/models"
	"e-commerce-api/service"
	"github.com/gofiber/fiber/v3"
)

func GetProduct(c Fiber.Ctx) error {
	product := service.GetAllProduct()
	return c.JSON(product)
}

func FindProduct(c Fiber.Ctx) error {
	id := c.Params("id")
	product := service.GetProductById(id)
	return c.JSON(product)
}

func CreateProduct(c Fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil{
		return c.Status(400).SendString(err.Error())
	}
	new_product:= service.CreateProduct(product)
	return c.JSON(new_product)
}


