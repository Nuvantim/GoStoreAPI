package service

import (
	"api/models"
	"api/database"
)

func GetCart(id string) models.Cart{
	var cart models.Cart
	database.DB.Preload("CartItem").Find(&id, ID)
	return cart
}
