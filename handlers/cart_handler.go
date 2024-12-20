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
	var product = service.FindProduct(cart.ProductID)
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
	// Ambil ID user dari context
	// user_id := c.Locals("id_user").(uint)

	// Bind body request ke struct cart
	var cart models.Cart
	if err := c.Bind().Body(&cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Cari cart berdasarkan ID
	carts := service.FindCart(cart.ID)

	// if carts.UserID != user_id {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "Unauthorized",
	// 	})
	// }
	product := service.FindProduct(carts.ProductID)

	if cart.Quantity > product.Stock {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Insufficient stock",
		})
	}

	// Hitung total biaya
	cost := cart.Quantity * product.Price

	// Update cart
	cart_update := service.UpdateCart(cart, cost)

	return c.Status(200).JSON(cart_update)

}

func DeleteCart(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	// user_id := c.Locals("id_user").(uint)
	// if carts.ID != user_id {
	// 	return c.Status(401).JSON(fiber.Map{
	// 		"error": "Unauthorized",
	// 	})
	// }
	service.DeleteCart(uint(id))
	return c.Status(200).JSON(fiber.Map{
		"message": "Success",
	})
}
