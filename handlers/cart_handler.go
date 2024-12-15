package handler

import (
	"api/models"
	"api/service"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

func GetCart(c fiber.Ctx) error {
	id := c.Locals("user_id").(uint)
	cart := service.GetCart(id)
	return c.Status(200).JSON(cart)
}

func FindCart(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	cart := service.FindCart(uint(id))
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
	if uint(product.Stock) < 1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Stock Product is empty",
		})
	}

	var cost uint = product.Price

	carts := service.AddCart(cart, id_user, cost)
	return c.Status(200).JSON(carts)
}

func UpdateCart(c fiber.Ctx) error {

	user_id := c.Locals("id_user").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	var cart models.Cart
	if err := c.Bind().Body(&cart); err != nil {
		return c.Status(400).JSON(err.Error)
	}
	//check cart id
	carts := service.FindCart(uint(id))
	if carts.ID == 0 {
		c.Status(404).JSON(fiber.Map{
			"message": "Cart Not Found",
		})
	}

	if carts.UserID != user_id {
		c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	//check product
	product := service.FindProduct(carts.ProductID)
	if product.ID == 0 {
		c.Status(404).JSON(fiber.Map{
			"message": "Product Not Found",
		})
	}

	// sum new cost
	var cost uint
	if uint(cart.Quantity) < uint(product.Stock) {
		cost = cart.Quantity * product.Price
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Stock is not Compatible",
		})
	}

	cart_update := service.UpdateCart(uint(id), cart, cost)
	return c.Status(200).JSON(cart_update)

}

func DeleteCart(c fiber.Ctx) error{
	id,_ := strconv.Atoi(c.Params("id"))
	user_id := c.Locals("id_user").(uint)
	carts := service.FindCart(uint(id))

	if uint(carts.ID) != user_id {
		return c.Status(401).JSON(fiber.Map{
			"error" : "Unauthorized",
		})
	}
	service.DeleteCart(uint(id))
	return c.Status(200).JSON(fiber.Map{
		"error": "Success",
	})
}
