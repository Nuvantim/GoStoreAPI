package service

import (
	"api/database"
	"api/models"
)

// get category
func GetAllCategory() []models.Category {
	var category []models.Category
	database.DB.Find(&category)
	return category
}

// get category from id
func GetCategoryById(id string) models.Category {
	var category models.Category
	database.DB.Find(&category, id)
	return category
}

// create category
func CreateCategory(category models.Category) models.Category {
	database.DB.Create(&category)
	return category
}

// update category
func UpdateCategory(id uint, category_request models.Category) models.Category {
	var category models.Category
	database.DB.First(&category, id)
	category.Name = category_request.Name
	database.DB.Save(&category)
	return category
}

// delete category
func DeleteCategory(id uint) error {
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return err
	}
	database.DB.Delete(&category)
	return nil
}
