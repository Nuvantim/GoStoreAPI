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
	database.DB.Preload("OrderItem").Where("user_id = ?", id).Find(&order)
	return order
}

func FindOrder(id uint) models.Order {
	var order models.Order
	database.DB.Preload("OrderItem").First(&order, id)
	return order
}

func CreateOrder(id_user, totalPrice uint) models.Order {
	order := models.Order{
		UserID: id_user,
		Total:  totalPrice,
	}
	database.DB.Create(&order)
	database.DB.Preload('OrderItem').First(&order, order.ID)
	return order
}

func DeleteOrder(id uint) error {
	// Memulai transaksi database
	tx := database.DB.Begin()

	// Hapus OrderItem berdasarkan OrderID
	if err := tx.Where("order_id = ?", id).Delete(&models.OrderItem{}).Error; err != nil {
		tx.Rollback() // Membatalkan transaksi jika ada error
		return err
	}

	// Hapus Order berdasarkan ID
	if err := tx.Delete(&models.Order{}, id).Error; err != nil {
		tx.Rollback() // Membatalkan transaksi jika ada error
		return err
	}

	// Commit transaksi jika semua berhasil
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

/*
SERVICE ORDER ITEM
*/
func CreateOrderItem(orderID uint, cartData []models.Cart) error {
	// Membuat slice untuk menyimpan data order items
	var orderItems []models.OrderItem

	// Menyiapkan data untuk batch insert
	for _, cart := range cartData {
		orderItems = append(orderItems, models.OrderItem{
			OrderID:    orderID,
			ProductID:  cart.ProductID,
			Quantity:   cart.Quantity,
			Total_Cost: cart.Total_Cost,
		})
	}

	// Batch insert ke database
	if err := database.DB.Create(&orderItems).Error; err != nil {
		return err
	}

	return nil
}
