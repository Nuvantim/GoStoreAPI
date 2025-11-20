package middleware

import (
	"api/pkg/guard"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"
)

// AuthAndRefreshMiddleware verifikasi token JWT menggunakan RS512
func AuthAndRefreshMiddleware(c *fiber.Ctx) error {
	var tokenString string
	authHeader := c.Get("Authorization")
	authCookie := c.Cookies("refresh_token")

	// Ambil token dari header Authorization
	if strings.HasPrefix(authHeader, "Bearer ") {
		tokenString = strings.TrimPrefix(authHeader, "Bearer ")
	}

	// Validasi access token
	if tokenString != "" {
		token, err := jwt.ParseWithClaims(tokenString, &guard.Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Pastikan metode signing adalah RS512
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || token.Method.Alg() != "RS512" {
				return nil, jwt.ErrSignatureInvalid
			}
			return guard.PublicKey, nil
		})

		// Jika access token valid, set user context dan lanjutkan
		if err == nil && token.Valid {
			if claims, ok := token.Claims.(*guard.Claims); ok {
				c.Locals("user_id", claims.UserID)
				c.Locals("email", claims.Email)
				c.Locals("roles", claims.Roles)
				c.Set("Authorization", authHeader)
				return c.Next()
			}
		} else {
			// Jika access token tidak valid, coba refresh token
			if authHeader != "" && authCookie != "" {
				refreshToken, err := jwt.ParseWithClaims(authCookie, &guard.RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
					// Pastikan metode signing adalah RS512
					if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || token.Method.Alg() != "RS512" {
						return nil, jwt.ErrSignatureInvalid
					}
					return guard.PublicKey, nil
				})

				if err == nil && refreshToken.Valid {
					if claims, ok := refreshToken.Claims.(*guard.RefreshClaims); ok {
						newAccessToken, err := guard.AutoRefreshToken(claims.UserID)
						if err == nil {
							// Validasi token baru
							token, err := jwt.ParseWithClaims(newAccessToken, &guard.Claims{}, func(token *jwt.Token) (interface{}, error) {
								if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok || token.Method.Alg() != "RS512" {
									return nil, jwt.ErrSignatureInvalid
								}
								return guard.PublicKey, nil
							})

							if err == nil && token.Valid {
								if claims, ok := token.Claims.(*guard.Claims); ok {
									c.Locals("user_id", claims.UserID)
									c.Locals("email", claims.Email)
									c.Locals("roles", claims.Roles)
									c.Set("Authorization", "Bearer "+newAccessToken)
									return c.Next()
								}
							} else {
								log.Printf("Error validating new access token: %v", err)
							}
						} else {
							log.Printf("Error refreshing access token: %v", err)
						}
					}
				} else {
					log.Printf("Refresh token invalid: %v", err)
				}
			}
		}
	}

	// Jika kedua token tidak valid, kembalikan response unauthorized
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Authentication required",
	})
}
