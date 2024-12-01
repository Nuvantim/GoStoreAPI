package service

import (
	"golang.org/x/crypto/bcrypt"
	"api/database"
	"api/models"
)

func RegisterUser(name, email, password, address string, phone uint) models.User {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashPassword),
		Address:  address,
		Phone:    phone,
	}
	database.DB.Create(&user)
	return user
}

func GetUser() []models.User {
	var user []models.User
	database.DB.Find(&user)
	return user
}

func FindUser(id string) models.User {
	var user models.User
	database.DB.Find(&user, id)
	return user
}

func UpdateUser(id, name, email, password, address string, phone uint) models.User {
	var user models.User
	database.DB.First(&user, id)
	
	user.Name = name
	user.Email = email
	user.Address = address
	user.Phone = phone
	if password != "" {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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
