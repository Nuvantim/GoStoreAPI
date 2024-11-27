package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func Setup() fiber.Handler {
	return func (c fiber.Ctx) error {
		//jwt auth
		if err := AuthAndRefreshMiddleware(c); err != nil{
			return err
		}
		//next
		return c.Next()
	}
}
