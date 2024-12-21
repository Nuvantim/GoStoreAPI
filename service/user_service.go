package service

import (
	"api/database"
	"api/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(users models.User) models.User {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)

	user := models.User{
		Name:     users.Name,
		Email:    users.Email,
		Password: string(hashPassword),
		Address:  users.Address,
		Phone:    users.Phone,
	}
	database.DB.Create(&user)
	database.DB.First(&user, user.ID)
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

func UpdateUser(id string, users models.User) models.User {
	var user models.User
	database.DB.First(&user, id)

	user.Name = users.Name
	user.Email = users.Email
	user.Address = users.Address
	user.Phone = users.Phone
	if users.Password != "" {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
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
