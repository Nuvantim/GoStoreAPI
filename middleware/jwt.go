package middleware

import (
	"api/utils"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

var jwtSecret = []byte(os.Getenv("API_KEY"))
var refreshSecret = []byte(os.Getenv("REFRESH_KEY"))

func AuthAndRefreshMiddleware(c fiber.Ctx) error {
	// Extract access token
	accessToken := extractToken(c)

	// Validate access token
	if accessToken != "" {
		claims, err := validateAccessToken(accessToken)
		if err == nil {
			setContextLocals(c, claims)
			return c.Next()
		}
	}

	// Attempt token refresh
	refreshToken := c.Cookies("refresh_token")
	if refreshToken != "" {
		newAccessToken, err := refreshAccessToken(refreshToken)
		if err == nil {
			setNewAccessTokenCookie(c, newAccessToken)
			return c.Next()
		}
	}

	// Unauthorized if no valid tokens
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Authentication required",
	})
}

func validateAccessToken(tokenString string) (*utils.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*utils.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func refreshAccessToken(refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &utils.RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*utils.RefreshClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	return utils.CreateToken(claims.UserID, claims.Email)
}

func extractToken(c fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return c.Cookies("access_token")
}

func setContextLocals(c fiber.Ctx, claims *utils.Claims) {
	c.Locals("user_id", claims.UserID)
	c.Locals("email", claims.Email)
}

func setNewAccessTokenCookie(c fiber.Ctx, newToken string) {
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    newToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
	})
}
