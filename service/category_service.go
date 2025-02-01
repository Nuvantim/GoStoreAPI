package service

import (
	"api/database"
	"api/models"
)

type Category models.Category // declare type model Category
var category Category // declare variabel Category

// get category
func GetAllCategory() []Category {
	var categorys []Category
	database.DB.Find(&categorys)
	return categorys
}

// get category from id
func FindCategory(id uint) Category {
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
	database.DB.Take(&category, id)
	category.Name = category_request.Name
	database.DB.Save(&category)
	return category
}

// delete category
func DeleteCategory(id uint) error {
	if err := database.DB.Take(&category, id).Error; err != nil {
		return err
	}
	database.DB.Delete(&category)
	return nil
}
