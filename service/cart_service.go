package service

import (
	"api/database"
	"api/models"
	"errors"
)

func GetCart(id uint) []models.Cart {
	var cart []models.Cart
	database.DB.Where("user_id = ?", id).Preload("Product").Preload("Product.Category").Find(&cart)
	return cart
}

func FindCart(id uint) models.Cart {
	var carts models.Cart
	database.DB.Where("id = ?", id).Preload("Product").Preload("Product.Category").Find(&carts)
	return carts
}

func AddCart(cart_data models.Cart, id_user, cost uint) models.Cart {
	cart := models.Cart{
		UserID:     id_user,
		ProductID:  cart_data.ProductID,
		Total_Cost: cost,
	}
	database.DB.Create(&cart)
	database.DB.Preload("Product").Preload("Product.Category").First(&cart, cart.ID)
	return cart
}

func UpdateCart(cart_update models.Cart, cost uint) models.Cart {
	var cart models.Cart
	// Ambil data cart berdasarkan ID
	database.DB.First(&cart, cart_update.ID)

	// Update menggunakan Updates() untuk langsung mengubah kolom yang diperlukan
	database.DB.Model(&cart).Preload("Product").Preload("Product.Category").Updates(map[string]interface{}{
		"Quantity":   cart_update.Quantity,
		"Total_Cost": cost,
	})

	return cart
}

func DeleteCart(input interface{}) error {
	switch v := input.(type) {
	case uint: // Jika input adalah ID tunggal
		return database.DB.Where("id = ?", v).Delete(&models.Cart{}).Error

	case []uint: // Jika input adalah array ID
		return database.DB.Where("id IN ?", v).Delete(&models.Cart{}).Error

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
