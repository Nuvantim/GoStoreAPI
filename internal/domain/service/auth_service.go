package service

import (
	"api/internal/database"
	"api/internal/domain/models"
	rds "api/internal/redis"
	"api/pkg/guard"
	"errors"
	"fmt"

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
	accessToken, err := guard.CreateToken(user.ID, user.Email, user.Roles)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := guard.CreateRefreshToken(user.ID, user.Email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func SendOTP(email string) (string, error) {
	otp, code := guard.GenerateOTP()

	token := models.Token{
		Otp:   otp,
		Email: email,
	}

	// set otp to redis
	if err := rds.SetData(fmt.Sprintf("verify:%d", token.Otp), token); err != nil {
		return "", err
	}

	if error := guard.SendOTP(email, code); error != nil {
		return "", error
	}

	// send OTP
	return "otp success send", nil

}

func UpdatePassword(email, password string) (string, error) {
	hash := guard.HashBycrypt(password)

	// find user by email
	var user models.User
	error := database.DB.Where("email = ?", email).Take(&user).Error
	if errors.Is(error, gorm.ErrRecordNotFound) {
		return "", errors.New("user not found")
	}
	user.Password = string(hash)
	database.DB.Save(&user)
	return "update password success", nil

}

func ValidateOTP(otp uint64) (*models.Token, error) {
	data, err := rds.GetData[models.Token](fmt.Sprintf("verify:%d", otp))
	if err != nil {
		return nil, fmt.Errorf("failed get otp: %s", err)
	}
	return data, nil
}
