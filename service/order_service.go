package service

import (
	"api/database"
	"api/models"
)
/*
SERVICE ORDER
*/
func GetOrder(id uint) []models.Order{
	var order []models.Order
	database.DB.Where("user_id = ?", id).Find(&order)
	return order
}

func FindOrder(id uint) models.Order{
	var order models.Order
	database.DB.First(&order, id)
	return order
}

func CreateOrder(id_user, totalPrice uint) models.Order {
	order := models.Order{
		UserID: id_user,
		Total:  totalPrice,
	}
	database.DB.Create(&order)
	database.DB.Preload("User").First(&order, order.ID)
	return order
}

func DeleteOrder(id uint) error{
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil{
		return rr
	}
	database.DB.Delete(&order)
}
