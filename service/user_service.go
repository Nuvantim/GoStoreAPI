package service

import (
	"api/database"
	"api/models"
	"api/utils"
)

func RegisterUser(users models.User) models.User {
	// hashing password
	hashPassword := utils.HashBycrypt(users.Password)

	user := models.User{
		Name:     users.Name,
		Email:    users.Email,
		Password: string(hashPassword),
	}
	database.DB.Create(&user)
	
	info := models.UserInfo{
		UserID : user.ID,
	}
	
	database.DB.Create(&info)
	alert := make[string]interface{}{
		"message" : "Success Register, Please Check Your Email"
	}
	return alert
}

func GetUser() []models.User {
	var user []models.User
	database.DB.Find(&user)
	return user
}

func FindUser(id string) models.User {
	var user models.User
	database.DB.First(&user, id)
	database.DB.Where("user_id = ?",user.ID).First(&info)

	data := make[string]interface{}{
		"user" : user,
		"user_info" : info,
	}
	return data
}

func UpdateUser(id string, users models.User) models.User {
	var user models.User
	database.DB.First(&user, id)

	user.Name = users.Name
	user.Email = users.Email
	user.Address = users.Address
	user.Phone = users.Phone
	if users.Password != "" {
		hashPassword := utils.HashBycrypt(users.Password)
		user.Password = string(hashPassword)
	}
	database.DB.Save(&user)
	return user
}

func DeleteUser(id string) error {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return err
	}
	database.DB.Delete(&user)
	return nil
}
