package service

import (
	"api/internal/database"
	"api/internal/domain/models"
)

type Category = models.Category // declare type model Category

// get category
func GetAllCategory() []Category {
	var category []Category // declare variabel Category
	database.DB.Find(&category)
	return category
}

// get category from id
func FindCategory(id uint) Category {
	var category Category // declare variabel Category
	database.DB.Find(&category, id)
	return category
}

// create category
func CreateCategory(category Category) Category {
	database.DB.Create(&category)
	return category
}

// update category
func UpdateCategory(id uint, category_request Category) Category {
	var category Category // declare variabel Category
	database.DB.Take(&category, id)
	category.Name = category_request.Name
	database.DB.Save(&category)
	return category
}

// delete category
func DeleteCategory(id uint) error {
	var category Category // declare variabel Category
	if err := database.DB.Take(&category, id).Error; err != nil {
		return err
	}
	database.DB.Delete(&category)
	return nil
}
