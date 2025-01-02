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

func UpdateUser(users map[string]interface{}, id string) map[string]interface{} {
	var user models.User
	// Ambil data user berdasarkan id
	database.DB.First(&user, id)

	// Update user
	if name, ok := users["name"].(string); ok {
		user.Name = name
	}
	if email, ok := users["email"].(string); ok {
		user.Email = email
	}
	if password, ok := users["password"].(string); ok && password != "" {
		hashPassword := utils.HashBycrypt(password)
		user.Password = string(hashPassword)
	}
	database.DB.Save(&user)

	// Update UserInfo
	var userInfo models.UserInfo
	database.DB.Where("user_id = ?", id).First(&userInfo)

	if age, ok := users["age"].(uint); ok {
		userInfo.Age = age
	}
	if phone, ok := users["phone"].(uint); ok {
		userInfo.Phone = phone
	}
	if district, ok := users["district"].(string); ok {
		userInfo.District = district
	}
	if city, ok := users["city"].(string); ok {
		userInfo.City = city
	}
	if state, ok := users["state"].(string); ok {
		userInfo.State = state
	}
	if country, ok := users["country"].(string); ok {
		userInfo.Country = country
	}

	database.DB.Save(&userInfo)

	// Buat peta untuk data yang ingin dikembalikan
	data := map[string]interface{}{
		"user":      user,
		"user_info": userInfo,
	}

	// Kembalikan data
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
