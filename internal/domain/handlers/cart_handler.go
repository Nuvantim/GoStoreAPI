package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils/responses"

	"github.com/gofiber/fiber/v2"
)

type Cart = service.Cart // declare type models Cart

/*
HANDLER GET CART
*/
func GetCart(c *fiber.Ctx) error {
	id := c.Locals("user_id").(uint64) // get user id
	cart := service.GetCart(id)
	return c.Status(200).JSON(response.Pass("success get cart", cart))
}

/*
HANDLER FIND CART
*/
func FindCart(c *fiber.Ctx) error {
	// get user id & cart id
	user_id := c.Locals("user_id").(uint64)
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}

	// service Find a Cart
	cart, _ := service.FindCart(uint64(id))

	// check cart exist
	if cart.ID == 0 {
		return c.Status(404).JSON(response.Error("failed find cart", "cart not found"))
	}

	// check cart user
	if cart.UserID != user_id {
		return c.Status(403).JSON(response.Error("failed get cart", "cart data is forbidden"))
	}
	return c.Status(200).JSON(response.Pass("success find cart", cart))
}

/*
HANDLER CREATE CART
*/
func CreateCart(c *fiber.Ctx) error {
	var cart Cart                               // declare variabel Cart
	id_user, ok := c.Locals("user_id").(uint64) // get user id
	if !ok {
		return c.Status(401).JSON(response.Error("failed get id", "invalid user id"))
	}

	// parser json
	if err := c.BodyParser(&cart); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}

	// check product
	var product = service.FindProduct(cart.ProductID)
	if product.ID == 0 {
		return c.Status(404).JSON(response.Error("failed get product", "product not found"))
	}

	// check product
	if uint64(product.Stock) < 1 {
		return c.Status(400).JSON(response.Error("failed get product", "stock product is empty"))
	}

	// data total_cost
	var cost uint64 = product.Price

	// Connect service
	carts := service.CreateCart(cart, id_user, cost)
	return c.Status(200).JSON(response.Pass("success create cart", carts))
}

/*
HANDLER UPDATE CART
*/
func UpdateCart(c *fiber.Ctx) error {
	var cart Cart                               // declare variabel Cart
	user_id, ok := c.Locals("user_id").(uint64) // get user id
	if !ok {
		return c.Status(401).JSON(response.Error("failed get user id", "invalid user id"))
	}
	// Bind body request ke struct cart
	if err := c.BodyParser(&cart); err != nil {
		return c.Status(500).JSON(response.Error("failed parser json", err.Error()))
	}

	// Find cart
	carts, _ := service.FindCart(cart.ID)

	// check cart user
	if carts.UserID != user_id {
		return c.Status(403).JSON(response.Error("failed get cart", "cart data is forbidden"))
	}

	// Find Product
	product := service.FindProduct(carts.ProductID)

	// check product exist
	if product.ID == 0 {
		return c.Status(404).JSON(response.Error("failed get product", "product not found"))
	}

	// Check product stock
	if cart.Quantity > product.Stock {
		return c.Status(400).JSON(response.Error("failed get product", "insufficient stock"))
	}

	// accumulate quantity with price product
	cost := cart.Quantity * product.Price

	// start servce
	cart_update := service.UpdateCart(cart, cost)

	return c.Status(200).JSON(response.Pass("success update cart", cart_update))

}

/*
HANDLER DELETE CART
*/
func DeleteCart(c *fiber.Ctx) error {
	// Get id endpoint & user id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	user_id := c.Locals("user_id").(uint64)

	// Find Cart
	carts, _ := service.FindCart(uint64(id))
	if carts.UserID != user_id {
		return c.Status(403).JSON(response.Error("failed get cart", "cart data is forbidden"))
	}
	// Delete cart
	service.DeleteCart(uint64(id))
	return c.Status(200).JSON(response.Pass("cart deleted", struct{}{}))
}
