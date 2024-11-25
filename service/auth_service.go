package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

// Secret key used for JWT signing

var jwtSecret = []byte(os.Getenv("API_KEY"))

// User credentials - Replace with database calls
var users = map[string]string{
	"krisnayoga319@gmail.com": "$2a$10$yqwSo7kcWqoEr7Xff4806urpVhTSSzjExdwKple.e7OaDtYZ2DMe", // bcrypt password hash for 'password'
}

// Claims for JWT
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Login authenticates the user and returns a JWT token
func Login(email, password string) (string, error) {
	// Check if email exists
	hashedPassword, exists := users[email]
	if !exists {
		return "", errors.New("invalid email or password")
	}

	// Compare the password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Create JWT token
	token, err := CreateToken(email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// CreateToken generates a JWT token for the authenticated user
func CreateToken(email string) (string, error) {
	claims := Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken parses and validates the JWT token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// Check if token is valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
