package service

import (
        "e-commerce-api/database"
        "e-commerce-api/models"
)

//get category
func GetAllProduct()models.Product {
        var product models.Product
        database.DBConn.Find(&product)
        return product
}

//get Product from id
func GetProductById(id uint)models.Product {
        var product models.Product
        database.DBConn.Find(&product, id)
        return product
}

//create Product
func CreateProduct(Product models.Product) models.Product {
      database.DBConn.Create(&product)  
}