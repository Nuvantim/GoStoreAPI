package handler

import(
	"github.com/gofiber/fiber/v3"
	"api/service"
)

func GetCart (c fiber.Ctx) error {
	id := c.Locals("user_id").(uint)
	cart := service.GetCart(id)
	return c.Status(200).JSON(cart)
}