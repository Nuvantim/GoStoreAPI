package service

import (
	"api/models"
	"api/database"
)

func GetCart(id uint) models.Cart{
	var cart models.Cart
	database.DB.Where("user_id = ?", id).Preload("CartItem").Find(&cart)
	return cart
}
