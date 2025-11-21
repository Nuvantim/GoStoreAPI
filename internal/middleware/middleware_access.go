package middleware

import (
	"api/internal/domain/models"
	"github.com/gofiber/fiber/v2"
)

func Role(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get claims from the context that has been created by jwt
		claims := c.Locals("roles").([]models.Role)

		// Check Wheter the user has the required role
		hasRole := false
		for _, role := range claims {
			if role.Name == requiredRole {
				hasRole = true
				break
			}
		}

		if hasRole != true {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Role Forbidden",
			})
		}

		return c.Next()
	}
}

func Permission(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get claims from the context that has been created by jwt
		claims := c.Locals("roles").([]models.Role)

		// Check Wheter the user has the necessary permissions
		hasPermission := false
		for _, role := range claims {
			for _, permission := range role.Permissions {
				if permission.Name == requiredPermission {
					hasPermission = true
					break
				}
			}
			if hasPermission {
				break
			}
		}

		if !hasPermission {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Permission Forbidden",
			})
		}

		return c.Next()
	}
}
