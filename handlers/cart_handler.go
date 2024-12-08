package handler

import (
	"api/database"
	"api/models"
	"strconv"
	"api/service"
	"github.com/gofiber/fiber/v3"
)

func GetCart(c fiber.Ctx) error {
	id := c.Locals("user_id").(int)
	cart := service.GetCart(id)
	return c.Status(200).JSON(cart)
}

func FindCart(c fiber.Ctx) error {
	id := c.Params("id")
	cart := service.FindCart(id)
	return c.Status(200).JSON(cart)
}

func AddCart(c fiber.Ctx) error {
	var cart models.Cart
	id_user := c.Locals("user_id").(int)

	if err := c.Bind().Body(&cart); err != nil {
		return c.Status(400).JSON(err.Error)
	}
	//check product
	var product models.Product
	if err := database.DB.First(&product, cart.ProductID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Product not found",
		})
	}
	//check product
	if product.Stock == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Stock Product is empty",
		})
	}

	cost := product.Price

	carts := service.AddCart(cart, id_user, cost)
	return c.Status(200).JSON(carts)
}

func UpdateCart(c fiber.ctx) error {
	var cart models.Cart
	id_cart := c.Params("id")
	//check cart id
	carts := service.FindCart(id_cart)
	//check product
	product := service.FindProduct(cart.ProductID)

	//body request
	if err := c.Bind().Body(&cart); err != nil {
		return c.Status(400).JSON(err.Error)
	}
	var cost int
	if carts.Quantity < product.Stock {
		cost = cart.Quantity * product.Price
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Stock is not compatible",
		})
	}

	service.UpdateCart(cart, cost, id_cart)
	return cart

}
