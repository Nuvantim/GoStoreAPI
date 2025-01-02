package service

import (
	"api/database"
	"api/models"
	"api/utils"
)

func RegisterUser(users models.User) models.User {
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

func FindUser(id string) models.User {
	var user models.User
	database.DB.First(&user, id)
	database.DB.Where("user_id = ?",user.ID).First(&info)

	data := make[string]interface{} {
		"user" : user,
		"user_info" : info,
	}

	return data
}

func UpdateUser(users {}interface, id string) models.User {
	var user models.User
	database.DB.First(&user, id)
	//update User
	user.Name = users.Name
	user.Email = users.Email
	if users.Password != "" {
		hashPassword := utils.HashBycrypt(users.Password)
		user.Password = string(hashPassword)
	}
	database.DB.Save(&user)
	// update UserInfo
	var user_info models.UserInfo
	database.DB.Where("user_id = ? ", id).First(&user_info)
	user_info.Age = users.age
	user_info.Phone = users.phone
	user_info.District = users.disctrict
	user_info.City  =  users.city
	user_info.State = users.state
	user_info.Country =  users.country
	database.DB.Save(&user_info)

	data := make[string]interface{}{
		"user" : user,
		"user_info" : user_info,
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
