package service

import (
	"api/database"
	"api/models"
)

var(
	Product models.Product
	Products []models.Product
)

// get category
func GetAllProduct() []models.Product {
	database.DB.Select("id", "name", "price", "stock", "category_id").Preload("Category").Find(&Products)
	return Products
}

// get Product from id
func FindProduct(id uint) models.Product {
	database.DB.Preload("Category").
	Preload("Review").
	Preload("Review.User").
	Take(&Product, id)
	return Product
}

// create Product
func CreateProduct(product models.Product) models.Product {
	database.DB.Create(&product)
	database.DB.Preload("Category").Take(&product, product.ID)
	return Product
}

// update category
func UpdateProduct(id string, product_request models.Product) models.Product {
	database.DB.Take(&Product, id)
	Product.Name = product_request.Name
	Product.Description = product_request.Description
	Product.Price = product_request.Price
	Product.Stock = product_request.Stock
	Product.CategoryID = product_request.CategoryID
	database.DB.Save(&Product)
	database.DB.Preload("Category").Take(&Product, Product.ID)
	return Product
}

// delete product
func DeleteProduct(id string) error {
	if err := database.DB.Take(&Product, id).Error; err != nil {
		return err
	}
	database.DB.Delete(&Product)
	return nil
}
