package service

import (
	"api/database"
	"api/models"
)

// declare variable category
var ( 
	Category models.Category
	Categorys []models.Category
)

// get category
func GetAllCategory() []models.Category {
	database.DB.Find(&Categorys)
	return Categorys
}

// get category from id
func FindCategory(id uint) models.Category {
	database.DB.Find(&Category, id)
	return Category
}

// create category
func CreateCategory(category models.Category) models.Category {
	database.DB.Create(&category)
	return category
}

// update category
func UpdateCategory(id uint, category_request models.Category) models.Category {
	database.DB.Take(&Category, id)
	Category.Name = category_request.Name
	database.DB.Save(&Category)
	return Category
}

// delete category
func DeleteCategory(id uint) error {
	if err := database.DB.Take(&Category, id).Error; err != nil {
		return err
	}
	database.DB.Delete(&Category)
	return nil
}
