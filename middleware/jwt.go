package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"toy-store-api/service"
)

var jwtSecret = []byte(os.Getenv("API_KEY"))
var refreshSecret = []byte(os.Getenv("REFRESH_KEY"))

// AuthAndRefreshMiddleware checks and refreshes tokens
func AuthAndRefreshMiddleware(c fiber.Ctx) error {
	// Get the access token from cookies
	accessToken := c.Cookies("access_token")
	refreshToken := c.Cookies("refresh_token")

	// If access token exists, validate it
	if accessToken != "" {
		token, err := jwt.ParseWithClaims(accessToken, &service.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err == nil && token.Valid {
			// If access token is valid, proceed
			return c.Next()
		}
	}

	// If access token is invalid or expired, check refresh token
	if refreshToken != "" {
		refreshTokenClaims := &service.Claims{}
		refreshTokenObj, err := jwt.ParseWithClaims(refreshToken, refreshTokenClaims, func(token *jwt.Token) (interface{}, error) {
			return refreshSecret, nil
		})

		if err == nil && refreshTokenObj.Valid {
			// If refresh token is valid, generate a new access token
			newAccessToken, err := service.CreateToken(refreshTokenClaims.Email)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate new access token")
			}

			// Update the access token in cookies
			c.Cookie(&fiber.Cookie{
				Name:     "access_token",
				Value:    newAccessToken,
				HTTPOnly: true,
				Secure:   true, // Set to true in production
				SameSite: "Strict",
				Path:     "/",
			})

			// Continue to the next handler
			return c.Next()
		}
	}

	// If both tokens are invalid, return unauthorized
	return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
}
