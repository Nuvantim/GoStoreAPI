package middleware

import (
	"api/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

var jwtSecret = []byte(os.Getenv("API_KEY"))
var refreshSecret = []byte(os.Getenv("REFRESH_KEY"))

func AuthAndRefreshMiddleware(c fiber.Ctx) error {
	// Ambil token dari Authorization Header
	authHeader := c.Get("Authorization")
	var accessToken string
	if strings.HasPrefix(authHeader, "Bearer ") {
		accessToken = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		// Jika tidak ada header Authorization, coba ambil dari cookie
		accessToken = c.Cookies("access_token")
	}

	// Cek validitas access token
	if accessToken != "" {
		ClaimToken := &utils.Claims{}
		token, err := jwt.ParseWithClaims(accessToken, ClaimToken, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err == nil && token.Valid {
			// Simpan User ID dan Email ke c.Locals
			claims := token.Claims.(*utils.Claims)
			c.Locals("user_id", claims.UserID)
			c.Locals("email", claims.Email)
			return c.Next()
		}
	}

	// Ambil refresh token dari cookie
	refreshToken := c.Cookies("refresh_token")
	if refreshToken != "" {
		refreshTokenClaims := &utils.RefreshClaims{}
		refreshTokenObj, err := jwt.ParseWithClaims(refreshToken, refreshTokenClaims, func(token *jwt.Token) (interface{}, error) {
			return refreshSecret, nil
		})

		if err == nil && refreshTokenObj.Valid {
			// Buat token baru
			newAccessToken, err := utils.CreateToken(refreshTokenClaims.UserID, refreshTokenClaims.Email)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate new access token"})
			}

			// Set cookie untuk token baru
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

	// Jika kedua token tidak valid, kembalikan unauthorized
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
}
