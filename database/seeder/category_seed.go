package seeder

import (
	"api/database"
	"api/models"
	"log"
)

func seed_Category() {
	var count int64
	if err := database.DB.Model(&models.Category{}).Count(&count).Error; err != nil {
		log.Println(err)
	}
	if count == 0 {
		category := []models.Category{
			{Name: "Category 1"},
			{Name: "Category 2"},
		}

		for _, category := range category {
			if err := database.DB.Create(&category).Error; err != nil {
				log.Println(err)
			}
		}
	}

}
