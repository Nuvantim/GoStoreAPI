package service

import (
	"api/database"
	"api/models"
)

func GetCart(id uint) []models.Cart {
	var cart []models.Cart
	database.DB.Where("user_id = ?", id).Preload("Product").Find(&cart)
	return cart
}

func FindCart(id uint) models.Cart {
	var cart models.Cart
	database.DB.Preload("Product").First(&cart, id)
	return cart
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

// func DeleteCart(id uint) error {
// 	var cart models.Cart
// 	database.DB.First(&cart, id)
// 	database.DB.Delete(&cart)
// 	return nil

// }
