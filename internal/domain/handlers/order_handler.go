package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// struct Request
var request struct {
	CartID []uint64 `json:"cart_id" validate:"required"`
}

/*
HANDLER GET ORDER
*/
func GetOrder(c *fiber.Ctx) error {
	// get user id
	id := c.Locals("user_id").(uint64)

	// connect to service
	order := service.GetOrder(id)
	return c.Status(200).JSON(response.Pass("success get order", order))
}

/*
HANDLER FIND ORDER
*/
func FindOrder(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint64)
	id_order := c.Params("id")
	id, err := uuid.Parse(id_order)
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get uuid", err.Error()))
	}

	// connect service
	order := service.FindOrder(id)

	// check order exist
	if order.ID == uuid.Nil {
		return c.Status(404).JSON(response.Error("failed find order", "order not found"))
	}

	// check user order
	if order.UserID != user_id {
		return c.Status(403).JSON(response.Error("failed find order", "order forbidden"))
	}
	return c.Status(200).JSON(response.Pass("success find order", order))
}

/*
HANDLER CREATE ORDER
*/
func CreateOrder(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint64) // get id user

	var totalPrice uint64 //declare variabel totalPrice
	var totalItem uint64  //declare variabel totalItem

	// Bind body request ke struct request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("failed parser json", err.Error()))
	}

	//get all cart
	_, cart := service.FindCart(request.CartID)

	//check cart
	if cart == nil {
		return c.Status(404).JSON(response.Error("failed find cart", "cart not found"))
	}
	//sum total_cost
	for _, cart := range cart {

		// check cart user
		if cart.UserID != user_id {
			return c.Status(403).JSON(response.Error("failed find cart", "cart forbidden"))
		}
		// sum total price
		totalPrice += cart.Total_Cost
	}
	// sum total item
	totalItem = uint64(len(cart))

	//connect service
	order := service.CreateOrder(uint64(user_id), totalItem, totalPrice, cart)

	// remove cart after create order
	if err := service.DeleteCart(request.CartID); err != nil {
		return c.Status(500).JSON(response.Error("failed delete cart", err.Error()))
	}

	return c.Status(200).JSON(response.Pass("success create order", order))
}

/*
HANDLER DELETE ORDER
*/
func DeleteOrder(c *fiber.Ctx) error {
	// get endpoint id & user_id
	id_order := c.Params("id")
	id, err := uuid.Parse(id_order)
	if err != nil {
		return c.Status(400).JSON(response.Error("failed get uuid", err.Error()))
	}
	user_id := c.Locals("user_id")

	// find Order
	order := service.FindOrder(id)

	// cek order exist
	if order.ID == uuid.Nil {
		return c.Status(404).JSON(response.Error("failed find order", "order not found"))
	}

	// cek user order
	if order.UserID != user_id {
		return c.Status(403).JSON(response.Error("failed find order", "order forbidden"))
	}

	// connect service
	if err := service.DeleteOrder(id); err != nil {
		return c.Status(500).JSON(response.Error("failed delete order", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("order deleted", struct{}{}))

}
