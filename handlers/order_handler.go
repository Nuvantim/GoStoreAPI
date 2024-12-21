package handler

import (
	"api/service"
	"github.com/gofiber/fiber/v3"
)

// struct Request
var request struct {
	CartID []uint `json:"cart_id"`
}

func GetOrder(c fiber.Ctx) error {
	return nil
}

func FindOrder(c fiber.Ctx) error {
	return nil
}

func CreateOrder(c fiber.Ctx) error {
	// get id user
	user_id := c.Locals("user_id").(uint)

	//declare variabel totalPrice
	var totalPrice uint

	//convert json body to Request
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	//get all cart
	cart := service.TransferCart(request.CartID)

	//sum total_cost
	for i, _ := range cart {
		totalPrice += cart[i].Total_Cost
	}

	//create Order
	order := service.CreateOrder(uint(user_id), totalPrice)

	return c.Status(200).JSON(order)
}

func DeleteOrder(c fiber.Ctx) error {
	return nil
}
