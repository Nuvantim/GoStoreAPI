package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
	"toy-store-api/database"
	"toy-store-api/models"
)

var jwtSecret = []byte(os.Getenv("API_KEY"))

// / Update model Claims to include UserID
type Claims struct {
	UserID uint   `json:"user_id"` // Tambahkan User ID
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// Perbarui fungsi Login
func Login(email, password string) (string, error) {
	// Find user in the database
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("Internal Server Error")
	} else if err != nil {
		return "", err
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Create JWT token with User ID
	token, err := CreateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func CreateToken(userID uint, email string) (string, error) {
	claims := Claims{
		UserID: userID, // Tambahkan User ID
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
