package service

import (
	"api/database"
	"api/models"
)

func GetCart(id uint) models.Cart {
	var cart models.Cart
	database.DB.Where("user_id = ?", id).Preload("Product").Find(&cart)
	return cart
}

func FindCart(id uint) models.Cart {
	var cart models.Cart
	database.DB.First(&cart, id)
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

func UpdateCart(cart_data models.Cart, cost uint) models.Cart {
	var cart models.Cart
	database.DB.First(&cart, cart_data.ID)
	cart.Quantity = cart_data.Quantity
	cart.Total_Cost = cost
	database.DB.Save(&cart)
	database.DB.Preload("Product").First(&cart, cart.ID)
	return cart

}
