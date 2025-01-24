package service

// database.DB.Where("email = ?", email).Take(&user)
import (
	"api/database"
	"api/models"
	"api/utils"
)

func CheckEmail(email string) bool {
	var user models.User
	var usertemp models.UserTemp
	// check User
	var UserCount int64
	database.DB.Where("email = ?", email).Take(&user).Count(&UserCount)
	if UserCount > 0 {
		return true
	}

	// check UserTemp
	var UserTempCount int64
	database.DB.Where("email = ?", email).Take(&usertemp).Count(&UserTempCount)
	if UserTempCount > 0 {
		return true
	}
	return false
}

func RegisterAccount(users models.User) map[string]interface{} {
	// hashing password
	hashPassword := utils.HashBycrypt(users.Password)
	// create Otp
	otp := utils.GenerateOTP()
	// Buat user baru
	usertemp := models.UserTemp{
		Otp:      otp,
		Name:     users.Name,
		Email:    users.Email,
		Password: string(hashPassword),
	}
	// Simpan user ke database
	database.DB.Create(&usertemp)

	// Buat notifikasi sukses
	alert := map[string]interface{}{
		"message": "Success Register, Please Check Your Email",
	}
	return alert
}

// func GetUser() []models.User {
// 	var user []models.User
// 	database.DB.Find(&user)
// 	return user
// }

func FindAccount(id uint) map[string]interface{} {
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

func UpdateAccount(users models.User, user_info models.UserInfo, user_id uint) map[string]interface{} {
	// Ambil data user berdasarkan id
	var user models.User

	if err := database.DB.First(&user, user_id).Error; err != nil {
		return map[string]interface{}{"error": "User not found"}
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
		return map[string]interface{}{"error": "Failed to update user"}
	}

	// Ambil data user_info berdasarkan user_id
	var userInfo models.UserInfo
	if err := database.DB.Where("user_id = ?", user_id).First(&userInfo).Error; err != nil {
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

func DeleteAccount(user_id uint) error {
	var user models.User
	if err := database.DB.First(&user, user_id).Error; err != nil {
		return err
	}
	database.DB.Delete(&user)
	return nil
}
