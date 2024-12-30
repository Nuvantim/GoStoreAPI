package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"os"
	"errors"
)
// get api key from .env
var jwtSecret = []byte(os.Getenv("API_KEY"))
var refreshSecret = []byte(os.Getenv("REFRESH_KEY"))

// Claims untuk Access Token
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Claims untuk Refresh Token
type RefreshClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Buat Access Token
func CreateToken(userID uint, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), // Access token berlaku 2 jam
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Buat Refresh Token
func CreateRefreshToken(userID uint, email string) (string, error) {
	claims := RefreshClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Refresh token berlaku 30 hari
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}

// Fungsi untuk memperbarui access token menggunakan refresh token
func RefreshAccessToken(refreshToken string) (string, error) {
	// Parse refresh token
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	// Ambil klaim dari refresh token
	claims, ok := token.Claims.(*RefreshClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	// Buat access token baru
	newAccessToken, err := CreateToken(claims.UserID, claims.Email)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
