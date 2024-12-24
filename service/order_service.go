package service

import (
	"api/database"
	"api/models"
)

/*
SERVICE ORDER
*/
func GetOrder(id uint) []models.Order {
	var order []models.Order
	database.DB.Preload("Order_Item").Where("user_id = ?", id).Find(&order)
	return order
}

func FindOrder(id uint) models.Order {
	var order models.Order
	database.DB.Preload("Order_Item").First(&order, id)
	return order
}

func CreateOrder(id_user, totalPrice uint) models.Order {
	order := models.Order{
		UserID: id_user,
		Total:  totalPrice,
	}
	database.DB.Create(&order)
	database.DB.First(&order, order.ID)
	return order
}

func DeleteOrder(id uint) error {
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		return err
	}
	database.DB.Delete(&order)
	return nil
}

func CreateOrderItem(order_id uint,cart_data []models.Cart) error{
	for _,cart_data :=  range cart_data {
		order_item := models.OrderItem{
			OrderID : order_id,
			ProductID : cart_data.ProductID,
			Quantity : cart_data.Quantity,
			Total_Cost : cart_data.Total_Cost,
		}
		database.DB.Create(&order_item)
	}
	return nil
}

func DeleteOrderItem(order_id uint) error{
	database.DB.Where("order_id IN ?", order_id).Delete(&models.OrderItem{})
	return nil
}
