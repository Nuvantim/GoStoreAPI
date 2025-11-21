package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	"github.com/gofiber/fiber/v2"
)

type Product = service.Product //declare type models Product

func GetProduct(c *fiber.Ctx) error {
	products := service.GetAllProduct()
	return c.Status(200).JSON(response.Pass("success get product", products))
}

func FindProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	product := service.FindProduct(uint64(id))
	if product.ID == 0 {
		return c.Status(404).JSON(response.Error("failed find product", "product not found"))
	}
	return c.Status(200).JSON(response.Pass("success find product", product))
}

func CreateProduct(c *fiber.Ctx) error {
	var product Product //declare variabel Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(product); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}

	// check category
	ctg := service.FindCategory(product.CategoryID)
	if ctg.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(response.Error("failed find category", "category not found"))
	}

	products := service.CreateProduct(product)
	return c.Status(200).JSON(response.Pass("success create product", products))
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id") // get Params ID
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	var product Product //declare variabel Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(product); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}

	products := service.UpdateProduct(uint64(id), product)
	return c.Status(200).JSON(response.Pass("success update product", products))
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	if err := service.DeleteProduct(uint64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete product", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("product deleted", struct{}{}))
}
