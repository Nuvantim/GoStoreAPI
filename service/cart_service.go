package service

import (
	"api/database"
	"api/models"
	"errors"
	"gorm.io/gorm"
)

func GetCart(id uint) []models.Cart {
	var cart []models.Cart
	database.DB.Where("user_id = ?", id).
		Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "price", "stock", "category_id")
		}).
		Preload("Product.Category").
		Find(&cart)
	return cart
}

func FindCart(id uint) models.Cart {
	var carts models.Cart
	database.DB.Where("id = ?", id).
		Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "price", "stock", "category_id")
		}).
		Preload("Product.Category").
		Take(&carts)
	return carts
}

func AddCart(cart_data models.Cart, id_user, cost uint) models.Cart {
	cart := models.Cart{
		UserID:     id_user,
		ProductID:  cart_data.ProductID,
		Total_Cost: cost,
	}
	database.DB.Create(&cart)
	cart = FindCart(cart.ID)
	return cart
}

func UpdateCart(cart_update models.Cart, cost uint) models.Cart {
	var cart models.Cart
	//Get cart by ID
	database.DB.Take(&cart, cart_update.ID)

	// Update data
	cart.Quantity = cart_update.Quantity
	cart.Total_Cost = cost
	database.DB.Save(&cart)

	// return data
	cart = FindCart(cart.ID)
	return cart
}

func DeleteCart(input interface{}) error {
	var cart models.Cart
	switch v := input.(type) {
	case uint: // Jika input adalah ID tunggal
		return database.DB.Where("id = ?", v).Delete(&cart).Error

	case []uint: // Jika input adalah array ID
		return database.DB.Where("id IN ?", v).Delete(&cart).Error

	default:
		return errors.New("invalid input type")
	}
}

func TransferCart(cart_id []uint) []models.Cart {
	var cart []models.Cart
	result := database.DB.Where("id IN ?", cart_id).Limit(1).Find(&cart)
	if result.RowsAffected == 0 {
		return nil // Mengembalikan nil jika tidak ada data
	}

	database.DB.Where("id IN ?", cart_id).Find(&cart)
	return cart
}
