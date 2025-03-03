package service

import (
	"api/internal/database"
	"api/internal/domain/models"
)

type (
	Permission = models.Permission
)

func GetPermission() []Permission {
	var permission []Permission
	database.DB.Find(&permission)
	return permission
}
func CreatePermission(permission Permission) Permission {
	database.DB.Create(&permission)
	return permission
}
func FindPermission(id uint) Permission {
	var permission Permission
	database.DB.Take(&permission, id)
	return permission
}
func UpdatePermission(id uint, permissions Permission) Permission {
	var permission Permission
	// get data permission
	database.DB.Take(&permission, id)
	// update data permission
	permission.Name = permissions.Name
	// save data permission
	database.DB.Save(&permission)
	return permission
}
func DeletePermission(id uint) (string, error) {
	var permission Permission
	if err := database.DB.Take(&permission, id).Error; err != nil {
		return "", err
	}
	database.DB.Delete(&permission)
	return "Permission success deleted", nil
}
