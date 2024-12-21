package service

import(
	"api/database"
	"api/models"
)

func CreateOrder(id_user, totalPrice uint) models.Order{
	order := models.Order{
		UserID : id_user,
		Total : totalPrice,
	}
	database.DB.Create(&order)
	database.DB.Preload("User").First(&order, order.ID)
	return order
}
