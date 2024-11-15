package service

import (
        "e-commerce-api/database"
        "e-commerce-api/models"
)

//get category
func GetAllCategory()models.Category {
        var category models.Category
        database.DBConn.Find(&category)
        return category
}

//get category from id
func GetCategoryById(id uint) models.Category {
        var category models.Category
        database.DBConn.Find(&category, id)
        return category
}

//create category
func CreateCategory(category models.Category) models.Category {
      database.DBConn.Create(&category)
      return category  
}

//update category
func UpdateCategory(id uint, category_data models.Category) models.Category {
        var category models.Category
        database.DBConn.First(&category,id)
        category.Name = category_data.Name
        database.DBConn.Save(&category)
        return &category
}

//delete category
func DeleteCategory(id uint) error {
        var category models.Category
        if err := database.DBConn.First(&category, id).Error; err != nil {
                return err
        }
        database.DBConn.Delete(&category)
        return nil
}