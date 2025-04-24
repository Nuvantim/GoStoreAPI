package service

import (
	"api/internal/database"
)

func GetClient() []User {
	var user []User
	database.DB.Preload("Roles").Find(&user)
	return user
}

func FindClient(id uint64) User {
	var user User
	database.DB.Preload("Roles").Preload("Roles.Permissions").Take(&user, id)
	return user
}

func UpdateClient(userID uint64, roleID []uint64) User {
	var user User
	// Find Client
	database.DB.Take(&user, userID)

	//find role
	var role []Role
	database.DB.Where("id IN ?", roleID).Find(&role)

	// Replace Role
	database.DB.Model(&user).Association("Roles").Clear()
	database.DB.Model(&user).Association("Roles").Append(&role)
	users := FindClient(user.ID)
	return users

}

func RemoveClient(id uint64) error {
	var user User
	// Find Client
	if err := database.DB.Take(&user, id).Error; err != nil {
		return err
	}
	// Delete relation role from pivot table
	database.DB.Model(&user).Association("Role").Clear()
	// Delete
	database.DB.Delete(&user)
	return nil

}
