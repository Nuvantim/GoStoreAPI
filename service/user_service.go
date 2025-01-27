package service

import (
	"api/database"
	"api/models"
	"api/utils"
)

func CheckEmail(email string) bool {
	var user models.User
	var usertemp models.UserTemp

	// Check User
	result := database.DB.Where("email = ?", email).Find(&user)
	if result.RowsAffected > 0 {
		return true
	}

	// Check UserTemp
	resultTemp := database.DB.Where("email = ?", email).Find(&usertemp)
	if resultTemp.RowsAffected > 0 {
		return true
	}

	return false
}

func RegisterAccount(users models.UserTemp) string {
	// hashing password
	hashPassword := utils.HashBycrypt(users.Password)

	// create Otp
	otp := utils.GenerateOTP()

	// Send OTP
	utils.SendOTP(users.Email, otp)

	// Create UserTemp
	usertemp := models.UserTemp{
		Otp:      otp,
		Name:     users.Name,
		Email:    users.Email,
		Password: string(hashPassword),
	}
	// Simpan user ke database
	database.DB.Create(&usertemp)

	return "Success Register, Please Check Your Email"
}


func FindAccount(id uint) map[string]interface{} {
	var user models.User
	var info models.UserInfo

	// Get data by ID
	database.DB.Take(&user, id)

	// Ambil data user_info berdasarkan user_id
	database.DB.Where("user_id = ?", user.ID).Take(&info)

	// Buat peta untuk data yang ingin dikembalikan
	data := map[string]interface{}{
		"user":      user,
		"user_info": info,
	}

	// Kembalikan data
	return data
}

func UpdateAccount(users models.User, user_info models.UserInfo, user_id uint) map[string]interface{} {
	// Ambil data user berdasarkan id
	var user models.User

	if err := database.DB.Take(&user, user_id).Error; err != nil {
		return map[string]interface{}{"error": "User not found"}
	}

	// update user
	user.Name = users.Name
	user.Email = users.Email
	if users.Password != "" {
		hash := utils.HashBycrypt(users.Password)
		user.Password = string(hash)
	}

	// Simpan perubahan user
	if err := database.DB.Save(&user).Error; err != nil {
		return map[string]interface{}{"error": "Failed to update user"}
	}

	// Ambil data user_info berdasarkan user_id
	var userInfo models.UserInfo
	if err := database.DB.Where("user_id = ?", user_id).Take(&userInfo).Error; err != nil {
		return map[string]interface{}{"error": "User info not found"}
	}

	// Update user_info
	userInfo.Age = user_info.Age
	userInfo.Phone = user_info.Phone
	userInfo.District = user_info.District
	userInfo.City = user_info.City
	userInfo.State = user_info.State
	userInfo.Country = user_info.Country

	// Simpan perubahan user_info
	if err := database.DB.Save(&userInfo).Error; err != nil {
		return map[string]interface{}{"error": "Failed to update user info"}
	}

	// Kembalikan data yang telah diperbarui
	data := map[string]interface{}{
		"user":      user,
		"user_info": userInfo,
	}

	return data
}

func DeleteAccount(user_id uint) error {
	var user models.User
	if err := database.DB.Take(&user, user_id).Error; err != nil {
		return err
	}
	database.DB.Delete(&user)
	return nil
}
