package service

import (
	"api/database"
	"api/models"
	"gorm.io/gorm"
)

type Product = models.Product 		// declare type models Product

// get category
func GetAllProduct() []Product {
	var product []Product
	database.DB.Select("id", "name", "price", "stock", "category_id").Preload("Category").Find(&product)
	return product
}

// get Product from id
func FindProduct(id uint) Product {
	var product Product 			//declare variabel Product
	database.DB.Preload("Category").
	Preload("Review").
	Preload("Review.User").
	Take(&product, id)
	return product
}

// create Product
func CreateProduct(product Product) Product {
	database.DB.Create(&product)
	database.DB.Preload("Category").Take(&product, product.ID)
	return product
}

// update category
func UpdateProduct(id uint, product_request Product) Product {
	var product Product         	//declare variabel Product
	database.DB.Take(&product, id)
	product.Name = product_request.Name
	product.Description = product_request.Description
	product.Price = product_request.Price
	product.Stock = product_request.Stock
	product.CategoryID = product_request.CategoryID
	database.DB.Save(&product)
	database.DB.Preload("Category").Take(&product, product.ID)
	return product
}

// delete product
func DeleteProduct(id uint) error {
	var product Product         	//declare variabel Product
	if err := database.DB.Take(&product, id).Error; err != nil {
		return err
	}
	database.DB.Delete(&product)
	return nil
}
