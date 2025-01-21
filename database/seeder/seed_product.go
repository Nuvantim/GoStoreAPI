package seeder

import (
	"api/database"
	"api/models"
	"log"
)

func seed_Product() {
	var count int64
	if err := database.DB.Model(&models.Product{}).Count(&count).Error; err != nil {
		log.Println(err)
	}
	if count == 0 {
		product := []models.Product{
			{Name: "Product 1", Description: "Cool Product", Price: 100000, Stock: 10, CategoryID: 1},
			{Name: "Product 2", Description: "Cool Product", Price: 200000, Stock: 20, CategoryID: 2},
		}

		for _, product := range product {
			if err := database.DB.Create(&product).Error; err != nil {
				log.Println(err)
			}
		}
	}
}
