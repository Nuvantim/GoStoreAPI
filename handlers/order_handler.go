package handler

import (
	"api/service"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

// struct Request
var request struct {
	CartID []uint `json:"cart_id" validate:"required"`
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
	id_order := c.Params("id")
	id, _ := uuid.Parse(id_order)

	// connect service
	order := service.FindOrder(id)

	// check order exist
	if order.ID == uuid.Nil {
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
	user_id := c.Locals("user_id").(uint) // get id user

	var totalPrice uint //declare variabel totalPrice
	var totalItem uint  //declare variabel totalItem

	// Bind body request ke struct request
	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	//get all cart
	_,cart := service.FindCart(request.CartID)
	//check cart
	if cart == nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cart not found",
		})
	}
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
	// sum total item
	totalItem = uint(len(cart))

	//connect service
	order := service.CreateOrder(uint(user_id), totalItem, totalPrice, cart)

	// remove cart after create order
	service.DeleteCart(request.CartID)

	return c.Status(200).JSON(order)
}

/*
HANDLER DELETE ORDER
*/
func DeleteOrder(c fiber.Ctx) error {
	// get endpoint id & user_id
	id_order := c.Params("id")
	id, _ := uuid.Parse(id_order)
	user_id := c.Locals("user_id")

	// find Order
	order := service.FindOrder(id)

	// cek order exist
	if order.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Order Not Found",
		})
	}

	// cek user order
	if order.UserID != user_id {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbiden",
		})
	}
	// connect service
	service.DeleteOrder(id)
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
	})

}
