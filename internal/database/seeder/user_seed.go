package seeder

import (
	"api/internal/database"
	"api/internal/domain/models"
	"api/pkg/utils"
	"log"
)

func seed_User() {
	var count int64
	if err := database.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		log.Println(err)
	}
	if count == 0 {
		password := utils.HashBycrypt("12345678")
		user := models.User{
			Name:     "Yoga",
			Email:    "yoga@gmail.com",
			Password: string(password),
		}

		if err := database.DB.Create(&user).Error; err != nil {
			log.Println(err)
		}
		// user info
		info := models.UserInfo{
			UserID: user.ID,
		}
		if err := database.DB.Create(&info).Error; err != nil {
			log.Println(err)
		}
	}
}
