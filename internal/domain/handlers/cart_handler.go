package handler

import (
	"api/internal/domain/service"
	"github.com/gofiber/fiber/v2"
)

type Cart = service.Cart // declare type models Cart

/*
HANDLER GET CART
*/
func GetCart(c *fiber.Ctx) error {
	id := c.Locals("user_id").(uint64) // get user id
	cart := service.GetCart(id)
	return c.Status(200).JSON(cart)
}

/*
HANDLER FIND CART
*/
func FindCart(c *fiber.Ctx) error {
	// get user id & cart id
	user_id := c.Locals("user_id").(uint64)
	id, _ := c.ParamsInt("id")

	// service Find a Cart
	cart, _ := service.FindCart(uint64(id))

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
func CreateCart(c *fiber.Ctx) error {
	var cart Cart                           // declare variabel Cart
	id_user := c.Locals("user_id").(uint64) // get user id

	// Bind body request ke struct cart
	if err := c.BodyParser(&cart); err != nil {
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
	if uint64(product.Stock) < 1 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Stock Product is empty",
		})
	}

	// data total_cost
	var cost uint64 = product.Price

	// Connect service
	carts := service.CreateCart(cart, id_user, cost)
	return c.Status(200).JSON(carts)
}

/*
HANDLER UPDATE CART
*/
func UpdateCart(c *fiber.Ctx) error {
	var cart Cart                           // declare variabel Cart
	user_id := c.Locals("user_id").(uint64) // get user id

	// Bind body request ke struct cart
	if err := c.BodyParser(&cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Find cart
	carts, _ := service.FindCart(cart.ID)

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
func DeleteCart(c *fiber.Ctx) error {
	// Get id endpoint & user id
	id, _ := c.ParamsInt("id")
	user_id := c.Locals("user_id").(uint64)

	// Find Cart
	carts, _ := service.FindCart(uint64(id))
	if carts.UserID != user_id {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}
	// Delete cart
	service.DeleteCart(uint64(id))
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
