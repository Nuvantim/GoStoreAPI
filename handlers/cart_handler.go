package handler

import(
	"api/models"
	"github.com/gofiber/fiber/v3"
)

func GetCart (c fiber.Ctx) error {
	id := c.Local("user_id")
	cart := service.GetCart(id)
	return c.Status(200).JSON(cart)
}