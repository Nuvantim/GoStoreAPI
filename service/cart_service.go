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

func AddCart(cart_data models.Cart, user_id, cost uint) models.Cart {
	cart := models.Cart{
		UserID:     user_id,
		ProductID:  cart_data.ProductID,
		Total_Cost: cost,
	}
	database.DB.Create(&cart)
	database.DB.Preload("Product").First(&cart, cart.ID)
	return cart
}

func UpdateCart(cart_update models.Cart, cost uint) models.Cart {
	var cart models.Cart
	database.DB.First(&cart, cart_update.ID)
	cart.Quantity = cart_update.Quantity
	cart.Total_Cost = cost

	database.DB.Update(&cart)

	return cart
}

// func DeleteCart(id uint) error {
// 	var cart models.Cart
// 	database.DB.First(&cart, id)
// 	database.DB.Delete(&cart)
// 	return nil

// }
