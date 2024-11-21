package service

import (
	"e-commerce-api/database"
	"e-commerce-api/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUser() []models.User {
	var user []models.User
	database.DBConn.Find(&user)
	return user
}

func FindUser(id string) models.User {
	var user models.User
	database.DBConn.Find(&user, id)
	return user
}

func RegisterUser(name, email, password, address string, phone uint) models.User {
	hashPassword,_:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashPassword),
		Address:  address,
		Phone:    phone,
	}
	database.DBConn.Create(&user)
	return user
}
