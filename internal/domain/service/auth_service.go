package service

import (
	"api/internal/database"
	"api/internal/domain/models"
	"api/pkg/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login Action
func Login(email, password string) (string, string, error) {
	// Find User in Database
	var user models.User
	err := database.DB.Where("email = ?", email).Take(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", errors.New("user not found")
	} else if err != nil {
		return "", "", err
	}

	// Compared Database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Create access token and refresh token
	accessToken, err := utils.CreateToken(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.CreateRefreshToken(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Verification Otp
func OtpVerify(otp string) string {
	// check otp
	if otp == "" {
		return "OTP cannot be empty"
	}

	// search user temporary
	var userTemp models.UserTemp
	result := database.DB.Where("otp = ?", otp).Take(&userTemp)
	if result.Error != nil {
		return "Invalid OTP"
	}

	// Start transaction database
	tx := database.DB.Begin()
	if tx.Error != nil {
		return "Failed to start database transaction"
	}

	// Create new User
	user := models.User{
		Name:     userTemp.Name,
		Email:    userTemp.Email,
		Password: userTemp.Password,
	}

	// Save database
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return "Failed to create user"
	}

	// Create user info
	userInfo := models.UserInfo{
		UserID: user.ID,
	}

	// save user info on database
	if err := tx.Create(&userInfo).Error; err != nil {
		tx.Rollback()
		return "Failed to create user infor"
	}

	// delete user temporary
	if err := tx.Delete(&userTemp).Error; err != nil {
		tx.Rollback()
		return "Failed to delete temporary user"
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return "Failed to commit transaction"
	}

	return "Verification successful, please login!"
}
