package service

import (
	"api/internal/database"
	"api/internal/domain/models"
	"api/pkg/utils"
	"errors"
)

type ( // declare type models User & UserInfo
	User     = models.User
	UserInfo = models.UserInfo
)

func CheckEmail(email string) uint {
	// declare variabel models
	var user User
	// declare count variabel
	var countUser int64

	database.DB.Model(&user).Where("email = ?", email).Count(&countUser)

	return uint(countUser)
}

func RegisterAccount(name, email, password string) (string, error) {
	// hash password
	hash := utils.HashBycrypt(password)
	// Create data user
	user := User{
		Name:     name,
		Email:    email,
		Password: string(hash),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return "", err
	}

	// create data user info
	userInfo := UserInfo{
		UserID: user.ID,
	}
	if err := database.DB.Create(&userInfo).Error; err != nil {
		return "", err
	}

	return "Your account successfuly registered", nil
}

func FindAccount(id uint) (User, UserInfo) {
	var user User
	var info UserInfo

	// Get data by ID
	database.DB.Take(&user, id)

	// Ambil data user_info berdasarkan user_id
	database.DB.Where("user_id = ?", user.ID).Take(&info)

	// Kembalikan data
	return user, info
}

func UpdateAccount(users User, user_info UserInfo, user_id uint) (User, UserInfo, error) {
	// Declare variable
	var user User
	var userInfo UserInfo

	// Get user data by id
	if err := database.DB.Take(&user, user_id).Error; err != nil {
		return user, userInfo, errors.New("User not found")
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
		return user, userInfo, errors.New("Failed to update user")
	}

	// Ambil data user_info berdasarkan user_id
	if err := database.DB.Where("user_id = ?", user_id).Take(&userInfo).Error; err != nil {
		return user, userInfo, errors.New("User info not found")
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
		return user, userInfo, errors.New("Failed to update user info")
	}

	return users, userInfo, nil
}

func DeleteAccount(user_id uint) error {
	var user User
	if err := database.DB.Take(&user, user_id).Error; err != nil {
		return err
	}
	database.DB.Delete(&user)
	return nil
}
