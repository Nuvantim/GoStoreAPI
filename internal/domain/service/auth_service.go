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
	err := database.DB.Where("email = ?", email).Preload("Roles").Preload("Roles.Permissions").Take(&user).Error
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
	accessToken, err := utils.CreateToken(user.ID, user.Email, user.Roles)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.CreateRefreshToken(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func SendOTP(email string) (string, error) {
	otp := utils.GenerateOTP()

	token := models.Token{
		Otp:   otp,
		Email: email,
	}
	if err := database.DB.Create(&token).Error; err != nil {
		return "", err
	}

	if error := utils.SendOTP(email, otp); error != nil {
		return "", error
	}

	// send OTP
	return "otp success send", nil

}

func UpdatePassword(otp uint, password string) (string, error) {
	var token models.Token
	// find otp
	err := database.DB.Where("otp = ?", otp).Take(&token).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("Otp code not found")
	}
	hash := utils.HashBycrypt(password)

	// find user by email
	var user models.User
	error := database.DB.Where("email = ?", token.Email).Take(&user).Error
	if errors.Is(error, gorm.ErrRecordNotFound) {
		return "", errors.New("User not found")
	}
	user.Password = string(hash)
	database.DB.Save(&user)
	return "Update password success", nil

}

func ValidateOTP(otp uint, email string) models.Token {
	var token models.Token
	database.DB.Where("email = ?", email).Where("otp = ?", otp).Take(&token)
	return token
}

func DeleteOTP(otp uint) error {
	var token models.Token
	if err := database.DB.Where("otp = ?", otp).Take(&token).Error; err != nil {
		return err
	}
	database.DB.Delete(&token)
	return nil
}
