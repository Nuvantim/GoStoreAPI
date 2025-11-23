package service

import (
	"api/internal/database"
	"api/internal/domain/models"
	"api/pkg/guard"
	"errors"
)

type ( // declare type models User & UserInfo
	User     = models.User
	UserInfo = models.UserInfo
)

func CheckEmail(email string) uint64 {
	// declare variabel models
	var user User
	// declare count variabel
	var countUser int64

	database.DB.Model(&user).Where("email = ?", email).Count(&countUser)

	return uint64(countUser)
}

func RegisterAccount(name, email, password string) (string, error) {
	// hash password
	hash := guard.HashBycrypt(password)
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

	return "your account successfuly registered", nil
}

func FindAccount(id uint64) (User, UserInfo) {
	var user User
	var info UserInfo

	// Get data by ID
	database.DB.Take(&user, id)

	// Ambil data user_info berdasarkan user_id
	database.DB.Where("user_id = ?", user.ID).Take(&info)

	// hide user id
	user.ID = 0

	// Kembalikan data
	return user, info
}

func UpdateAccount(users User, user_info UserInfo, user_id uint64) (User, UserInfo, error) {
	// Declare variable
	var user User
	var userInfo UserInfo

	// Get user data by id
	if err := database.DB.Take(&user, user_id).Error; err != nil {
		return user, userInfo, errors.New("user not found")
	}

	// update user
	user.Name = users.Name
	user.Email = users.Email
	if users.Password != "" {
		hash := guard.HashBycrypt(users.Password)
		user.Password = string(hash)
	}

	// Simpan perubahan user
	if err := database.DB.Save(&user).Error; err != nil {
		return user, userInfo, errors.New("failed to update user")
	}

	// Ambil data user_info berdasarkan user_id
	if err := database.DB.Where("user_id = ?", user_id).Take(&userInfo).Error; err != nil {
		return user, userInfo, errors.New("user info not found")
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
		return user, userInfo, errors.New("failed to update user info")
	}

	return users, userInfo, nil
}

func DeleteAccount(user_id uint64) error {
	var user User
	if err := database.DB.Take(&user, user_id).Error; err != nil {
		return err
	}
	database.DB.Delete(&user)
	return nil
}
