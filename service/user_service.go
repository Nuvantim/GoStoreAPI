package service

import (
	"golang.org/x/crypto/bcrypt"
	"toy-store-api/database"
	"toy-store-api/models"
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
	database.DBConn.Create(&user)
	return user
}

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

func UpdateUser(id, name, email, password, address string, phone uint) models.User {
	var user models.User
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	database.DBConn.First(&user, id)
	//declare data
	users := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashPassword),
		Address:  address,
		Phone:    phone,
	}
	//put
	user.Name = users.Name
	user.Email = users.Email
	user.Password = users.Password
	user.Address = users.Address
	user.Phone = users.Phone
	database.DBConn.Save(&user)
	return user
}

func DeleteUser(id string) error {
	var user models.User
	if err := database.DBConn.First(&user, id).Error; err != nil {
		return err
	}
	database.DBConn.Delete(&user)
	return nil
}
