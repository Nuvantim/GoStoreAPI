package handler

import (
	"api/service"
	"github.com/gofiber/fiber/v3"
	"strconv"
)
var cart service.Cart // declare variabel Cart
/*
HANDLER GET CART
*/
func GetCart(c fiber.Ctx) error {
	// get user id
	id := c.Locals("user_id").(uint)

	// connect to service
	cart := service.GetCart(id)
	return c.Status(200).JSON(cart)
}

/*
HANDLER FIND CART
*/
func FindCart(c fiber.Ctx) error {
	// get user id
	user_id := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))
	cart := service.FindCart(uint(id))

	// check cart exist
	if cart.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cart Not Found",
		})
	}

	// check cart user
	if cart.UserID != user_id {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}
	return c.Status(200).JSON(cart)
}

/*
HANDLER CREATE CART
*/
func AddCart(c fiber.Ctx) error {
	// get user id
	id_user := c.Locals("user_id").(uint)

	// Bind body request ke struct cart
	if err := c.Bind().Body(&cart); err != nil {
		return c.Status(400).JSON(err.Error)
	}

	// check product
	var product = service.FindProduct(cart.ProductID)
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product Not Found",
		})
	}

	// check product
	if uint(product.Stock) < 1 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Stock Product is empty",
		})
	}

	// data total_cost
	var cost uint = product.Price

	// Connect service
	carts := service.CreateCart(cart, id_user, cost)
	return c.Status(200).JSON(carts)
}

/*
HANDLER UPDATE CART
*/
func UpdateCart(c fiber.Ctx) error {
	// get user id
	user_id := c.Locals("user_id").(uint)

	// Bind body request ke struct cart
	if err := c.Bind().Body(&cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Find cart
	carts := service.FindCart(cart.ID)

	// check cart user
	if carts.UserID != user_id {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	// Find Product
	product := service.FindProduct(carts.ProductID)

	// check product exist
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product Not Found",
		})
	}

	// Check product stock
	if cart.Quantity > product.Stock {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Insufficient stock",
		})
	}

	// accumulate quantity with price product
	cost := cart.Quantity * product.Price

	// Update cart
	cart_update := service.UpdateCart(cart, cost)

	return c.Status(200).JSON(cart_update)

}

/*
HANDLER DELETE CART
*/
func DeleteCart(c fiber.Ctx) error {
	// Get id endpoint & user id
	id, _ := strconv.Atoi(c.Params("id"))
	user_id := c.Locals("user_id").(uint)

	// Find Cart
	carts := service.FindCart(uint(id))
	if carts.UserID != user_id {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}
	// Delete cart
	service.DeleteCart(uint(id))
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
