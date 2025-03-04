package seeder

import (
	"api/internal/database"
	"api/internal/domain/models"
	"api/internal/domain/service"
	"log"
)

func seed_Access() {
	// permission
	var countPermission int64
	if err := database.DB.Model(&models.Permission{}).Count(&countPermission).Error; err != nil {
		log.Println(err)
	}
	if countPermission == 0 {
		permission := []models.Permission{
			{Name: "kelola category"},
			{Name: "kelola product"},
		}
		// create permission
		database.DB.Create(&permission)
	}
	// role
	var countRole int64
	if err := database.DB.Model(&models.Role{}).Count(&countRole).Error; err != nil {
		log.Println(err)
	}
	if countRole == 0 {
		service.CreateRole("admin", []uint{})
		service.CreateRole("crew", []uint{1, 2})

		// assign role to client
		service.UpdateClient(1, []uint{1, 2})
	}
}
