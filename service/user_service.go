package service

import (
	"api/database"
	"api/models"
	"api/utils"
)

func RegisterUser(users models.User) map[string]interface{} {
	// hashing password
	hashPassword := utils.HashBycrypt(users.Password)

	// Buat user baru
	user := models.User{
		Name:     users.Name,
		Email:    users.Email,
		Password: string(hashPassword),
	}
	// Simpan user ke database
	database.DB.Create(&user)

	// Buat info user
	info := models.UserInfo{
		UserID: user.ID,
	}
	// Simpan info user ke database
	database.DB.Create(&info)

	// Buat notifikasi sukses
	alert := map[string]interface{}{
		"message": "Success Register, Please Check Your Email",
	}
	return alert
}


func GetUser() []models.User {
	var user []models.User
	database.DB.Find(&user)
	return user
}

func FindUser(id uint) map[string]interface{} {
	var user models.User
	var info models.UserInfo

	// Ambil data user berdasarkan id
	database.DB.First(&user, id)

	// Ambil data user_info berdasarkan user_id
	database.DB.Where("user_id = ?", user.ID).First(&info)

	// Buat peta untuk data yang ingin dikembalikan
	data := map[string]interface{}{
		"user":      user,
		"user_info": info,
	}

	// Kembalikan data
	return data
}

func UpdateUser(users models.User, user_info models.UserInfo, id uint) map[string]interface{} {
	// Ambil data user berdasarkan id
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return map[string]interface{}{"error": "User not found"}
	}

	// update user
	user.Name = users.Name
	user.Email = users.Email
	if users.Password != nil {
		hash := utils.HashBycrypt(users.Password)
		user.Password = string(hash)
	}
	
	// Simpan perubahan user
	if err := database.DB.Save(&user).Error; err != nil {
		return map[string]interface{}{"error": "Failed to update user"}
	}

	// Ambil data user_info berdasarkan user_id
	var userInfo models.UserInfo
	if err := database.DB.Where("user_id = ?", id).First(&userInfo).Error; err != nil {
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

func DeleteUser(id string) error {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return err
	}
	database.DB.Delete(&user)
	return nil
}
