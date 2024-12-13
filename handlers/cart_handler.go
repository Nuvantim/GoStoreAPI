package handler

import (
	"api/models"
	"api/service"
	"github.com/gofiber/fiber/v3"
)

func GetCart(c fiber.Ctx) error {
	id := c.Locals("user_id").(uint)
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
	id_user := c.Locals("user_id").(uint)

	if err := c.Bind().Body(&cart); err != nil {
		return c.Status(400).JSON(err.Error)
	}
	//check product
	var product = service.FindProduct(uint(cart.ProductID))
	//check product
	if uint (product.Stock) < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Stock Product is empty",
		})
	}

	var cost uint= product.Price

	carts := service.AddCart(cart, id_user, cost)
	return c.Status(200).JSON(carts)
}

func UpdateCart(c fiber.Ctx) error {

	var user_id = c.Locals("id_user")

	var cart models.Cart
	if err := c.Bind().Body(&cart); err != nil {
		return c.Status(400).JSON(err.Error)
	}
	//check cart id
	carts := service.FindCart(string(cart.ID))
	// if err != nil {
	// 	c.Status(404).JSON(fiber.Map{
	// 		"message": "Cart Not Found",
	// 	})
	// }
	if carts.UserID != user_id {
		c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	//check product
	product := service.FindProduct(uint(carts.ProductID))
	// if err != nil {
	// 	c.Status(404).JSON(fiber.Map{
	// 		"message": "Product Not Found",
	// 	})
	// }

	var cost uint
	if cart.Quantity < product.Stock {
		cost = cart.Quantity * product.Price
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Stock is not compatible",
		})
	}
	
	cart_update := service.UpdateCart(cart, cost)
	return c.Status(200).JSON(cart_update)

}
