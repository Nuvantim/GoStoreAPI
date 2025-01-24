package service

import (
	"api/database"
	"api/models"
	"api/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Fungsi Login
func Login(email, password string) (string, string, error) {
	// Cari user di database
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", errors.New("user not found")
	} else if err != nil {
		return "", "", err
	}

	// Bandingkan password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// Buat access token dan refresh token
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

func OtpVerify(otp string) map[string]interface{}{
	// check otp
	var usertemp models.UserTemp
	database.DB.Where("otp = ?", otp).Take(&usertemp)
	// Buat user baru
	user := models.User{
		Name:     usertemp.Name,
		Email:    usertemp.Email,
		Password: usertemp.Password,
	}
	// Simpan user ke database
	database.DB.Create(&user)

	// Buat info user
	info := models.UserInfo{
		UserID: user.ID,
	}
	// Simpan info user ke database
	database.DB.Create(&info)

	// delete userTemp
	var user_temp models.UserTemp
	database.DB.First(&user_temp, usertemp.ID)
	database.DB.Delete(&user_temp)

	alert := map[string]interface{}{
		"message": "Success Verification, Please Login !",
	}
	return alert
}
