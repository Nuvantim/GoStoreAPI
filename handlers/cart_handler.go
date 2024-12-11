package handler

import (
	"api/models"
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
	var product = service.FindProduct(string(cart.ProductID))
	//check product
	if int(product.Stock) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Stock Product is empty",
		})
	}

	cost := product.Price

	carts := service.AddCart(cart, id_user, cost)
	return c.Status(200).JSON(carts)
}

func UpdateCart(c fiber.Ctx) error {

	var id_cart = c.Params("id")
	var user_id = c.Locals("id_user")

	//check cart id
	carts := service.FindCart(id_cart)
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
	product := service.FindProduct(string(carts.ProductID))
	// if err != nil {
	// 	c.Status(404).JSON(fiber.Map{
	// 		"message": "Product Not Found",
	// 	})
	// }

	//body request
	var cart models.Cart
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
	
	cart_update := service.UpdateCart(cart, cost, cart_ID)
	return c.Status(200).JSON(cart_update)

}
