package handler

import (
	"api/service"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// struct Request
var request struct {
	CartID []uint `json:"cart_id"`
}

/*
HANDLER GET ORDER
*/
func GetOrder(c fiber.Ctx) error {
	// get user id
	id := c.Locals("user_id").(uint)

	// connect to service
	cart := service.GetOrder(id)
	return c.Status(200).JSON(cart)
}

/*
HANDLER FIND ORDER
*/
func FindOrder(c fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	// connect service
	order := service.FindOrder(uint(id))

	// check order exist
	if order.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Order Not Found",
		})
	}

	// check user order
	if order.UserID != user_id {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}
	return c.Status(200).JSON(order)
}

/*
HANDLER CREATE ORDER
*/
func CreateOrder(c fiber.Ctx) error {
	// get id user
	user_id := c.Locals("user_id").(uint)

	//declare variabel totalPrice
	var totalPrice uint

	// Bind body request ke struct request
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}


	//get all cart
	cart := service.TransferCart(request.CartID)

	//sum total_cost
	for _, cart := range cart {

		// check cart user
		if cart.UserID != user_id {
			return c.Status(403).JSON(fiber.Map{
				"message": "Forbidden",
			})
		}
		// sum total price
		totalPrice += cart.Total_Cost
	}

	//connect service
	order := service.CreateOrder(uint(user_id), totalPrice)
	service.CreateOrderItem(order.ID, cart)

	// remove cart after create order
	service.DeleteCart(request.CartID)

	return c.Status(200).JSON(order)
}

/*
HANDLER DELETE ORDER
*/
func DeleteOrder(c fiber.Ctx) error {
	// get endpoint id & user_id
	id, _ := strconv.Atoi(c.Params("id"))
	user_id := c.Locals("user_id").(uint)

	// find Order
	order := service.FindOrder(uint(id))

	// cek order exist
	if order.ID == 0 {
		c.Status(404).JSON(fiber.Map{
			"message": "Order Not Found",
		})
	}

	// cek user order
	if order.UserID != user_id {
		c.Status(403).JSON(fiber.Map{
			"message": "Forbiden",
		})
	}
	// connect service
	service.DeleteOrderItem(uint(id))
	service.DeleteOrder(uint(id))
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
	})

}
