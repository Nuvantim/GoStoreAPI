package handler

import (
	// "api/models"
	"api/service"
	"github.com/gofiber/fiber/v3"
)

func GetOrder(c fiber.Ctx) error {
  return nil
}

func FindOrder(c fiber.Ctx) error {
  return nil
}

func CreateOrder(c fiber.Ctx) error {
	// var carts []models.Cart
	var totalPrice uint

	var request struct {
		CartIDs []uint `json:"cart_id"`
	}

	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	for data,_ := range request {
		cart_data := service.FindCart(request[data].CartID)
		totalPrice += cart_data.Total_Cost
	}
	return c.Status(200).JSON(totalPrice)
}

func DeleteOrder (c fiber.Ctx) error {
  return nil
}
