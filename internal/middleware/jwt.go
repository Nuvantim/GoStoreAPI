package middleware

import (
	"api/pkg/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

var (
	jwtSecret     = []byte(os.Getenv("API_KEY"))
	refreshSecret = []byte(os.Getenv("REFRESH_KEY"))
)

func AuthAndRefreshMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenString := ""

	// Retrieve the token from the Authorization header or cookies
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		tokenString = c.Cookies("access_token")
	}

	// Try to validate the access token
	if tokenString != "" {
		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// If the access token is valid, set user context and proceed
		if err == nil && token.Valid {
			if claims, ok := token.Claims.(*utils.Claims); ok {
				c.Locals("user_id", claims.UserID)
				c.Locals("email", claims.Email)
				return c.Next()
			}
		}
	}

	// If access token validation fails, try to refresh it using the refresh token
	refreshToken := c.Cookies("refresh_token")
	if refreshToken != "" {
		token, err := jwt.ParseWithClaims(refreshToken, &utils.RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
			return refreshSecret, nil
		})

		// If the refresh token is valid, generate a new access token
		if err == nil && token.Valid {
			if claims, ok := token.Claims.(*utils.RefreshClaims); ok {
				newAccessToken, err := utils.CreateToken(claims.UserID, claims.Email)
				if err == nil {
					// Set the new access token in cookies
					c.Cookie(&fiber.Cookie{
						Name:     "access_token",
						Value:    newAccessToken,
						HTTPOnly: true,
						Secure:   true,
						SameSite: "Strict",
						Path:     "/",
					})

					// Set user context and proceed
					c.Locals("user_id", claims.UserID)
					c.Locals("email", claims.Email)

					return c.Next()
				}
			}
		}
	}

	// If both tokens are invalid, return an unauthorized response
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Authentication required",
	})
}
