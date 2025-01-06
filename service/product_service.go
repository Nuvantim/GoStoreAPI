package service

import (
	"api/database"
	"api/models"
)

// get category
func GetAllProduct() []models.Product {
	var product []models.Product
	database.DB.Select("id", "name", "price", "stock","category_id").Preload("Category").Find(&product)
	return product
}

// get Product from id
func FindProduct(id uint) models.Product {
	var product models.Product
	database.DB.Preload("Category").
		Preload("Review").
		Preload("Review.User").
		First(&product, id)
	return product
}

// create Product
func CreateProduct(product models.Product) models.Product {
	database.DB.Create(&product)
	database.DB.Preload("Category").First(&product, product.ID)
	return product
}

// update category
func UpdateProduct(id string, product_request models.Product) models.Product {
	var product models.Product
	database.DB.First(&product, id)
	product.Name = product_request.Name
	product.Description = product_request.Description
	product.Price = product_request.Price
	product.Stock = product_request.Stock
	product.CategoryID = product_request.CategoryID
	database.DB.Save(&product)
	database.DB.Preload("Category").First(&product, product.ID)
	return product
}

// delete product
func DeleteProduct(id string) error {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return err
	}
	database.DB.Delete(&product)
	return nil
}
