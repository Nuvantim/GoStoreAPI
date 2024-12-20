package service

import (
	"api/database"
	"api/models"
	"fmt"
)

func GetCart(id uint) []models.Cart {
	var cart []models.Cart
	database.DB.Where("user_id = ?", id).Preload("Product").Find(&cart)
	return cart
}

func FindCart(input interface{}) []models.Cart {
	var carts []models.Cart

	switch v := input.(type) {
	case uint: // Jika input adalah ID tunggal
		db.Where("id = ?", v).Find(&carts)
		return carts

	case []uint: // Jika input adalah array ID
		err := db.Where("id IN ?", v).Find(&carts)
		return carts

	default:
		return nil, fmt.Errorf("invalid input type")
	}
}


func UpdateCart(cart_update models.Cart, cost uint) models.Cart {
	var cart models.Cart
	// Ambil data cart berdasarkan ID
	database.DB.First(&cart, cart_update.ID)
	
	// Update menggunakan Updates() untuk langsung mengubah kolom yang diperlukan
	database.DB.Model(&cart).Updates(map[string]interface{}{
		"Quantity":   cart_update.Quantity,
		"Total_Cost": cost,
	})

	return cart
}

 func DeleteCart(input interface{}) error {
	switch v := input.(type) {
	case uint: // Jika input adalah ID tunggal
		return db.Where("id = ?", v).Delete(&models.Cart{}).Error

	case []uint: // Jika input adalah array ID
		return db.Where("id IN ?", v).Delete(&models.Cart{}).Error

	default:
		return fmt.Errorf("invalid input type")
	}
}

