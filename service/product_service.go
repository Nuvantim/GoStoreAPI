package service

import (
	"e-commerce-api/database"
	"e-commerce-api/models"
)

// get category
func GetAllProduct() []models.Product {
	var product []models.Product
	database.DBConn.Preload("Category").Find(&product)
	return product
}

// get Product from id
func GetProductById(id string) models.Product {
	var product models.Product
	database.DBConn.Find(&product, id)
	return product
}

// create Product
func CreateProduct(product models.Product) models.Product {
	database.DBConn.Create(&product)
	database.DBConn.Preload("Category").First(&product, product.ID)
	return product
}

// update category
func UpdateProduct(id string, product_request models.Product) models.Product {
	var product models.Product
	database.DBConn.First(&product, id)
	product.Name = product_request.Name
	database.DBConn.Save(&product)
	return product
}

// delete product
func DeleteProduct(id string) error {
	var product models.Product
	if err := database.DBConn.First(&product, id).Error; err != nil {
		return err
	}
	database.DBConn.Delete(&product)
	return nil
}
