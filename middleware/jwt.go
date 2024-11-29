package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"toy-store-api/service"
)

var jwtSecret = []byte(os.Getenv("API_KEY"))
var refreshSecret = []byte(os.Getenv("REFRESH_KEY"))

func AuthAndRefreshMiddleware(c fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	refreshToken := c.Cookies("refresh_token")

	if accessToken != "" {
		token, err := jwt.ParseWithClaims(accessToken, &service.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err == nil && token.Valid {
			// Simpan User ID dan Email ke c.Locals
			claims := token.Claims.(*service.Claims)
			c.Locals("user_id", claims.UserID)
			c.Locals("email", claims.Email)
			return c.Next()
		}
	}

	if refreshToken != "" {
		refreshTokenClaims := &service.Claims{}
		refreshTokenObj, err := jwt.ParseWithClaims(refreshToken, refreshTokenClaims, func(token *jwt.Token) (interface{}, error) {
			return refreshSecret, nil
		})

		if err == nil && refreshTokenObj.Valid {
			newAccessToken, err := service.CreateToken(refreshTokenClaims.UserID, refreshTokenClaims.Email)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate new access token")
			}

			c.Cookie(&fiber.Cookie{
				Name:     "access_token",
				Value:    newAccessToken,
				HTTPOnly: true,
				Secure:   true,
				SameSite: "Strict",
				Path:     "/",
			})

			// Simpan User ID dan Email ke c.Locals
			c.Locals("user_id", refreshTokenClaims.UserID)
			c.Locals("email", refreshTokenClaims.Email)
			return c.Next()
		}
	}

	// Jika kedua token tidak valid, log unauthorized
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}
