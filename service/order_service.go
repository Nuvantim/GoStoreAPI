package service

import (
	"api/database"
	"api/models"
	"github.com/google/uuid"
)

/*
SERVICE ORDER
*/
func GetOrder(id uint) []models.Order {
	var order []models.Order
	database.DB.Where("user_id = ?", id).Find(&order)
	return order
}

func FindOrder(id uuid.UUID) models.Order {
	var order models.Order
	database.DB.Preload("OrderItem.Product").First(&order,"id = ?", id)
	return order
}

func CreateOrder(id_user, totalItem, totalPrice uint, cartData []models.Cart) models.Order {
	// Create Order
	order := models.Order{
		UserID:     id_user,
		TotalPrice: totalPrice,
		TotalItem:  totalItem,
	}
	database.DB.Create(&order)

	// Create Order Item
	var orderItems []models.OrderItem

	// Menyiapkan data untuk batch insert
	for _, cart := range cartData {
		orderItems = append(orderItems, models.OrderItem{
			OrderID:   order.ID,
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			TotalCost: cart.Total_Cost,
		})
	}

	// Batch insert ke database
	database.DB.Create(&orderItems)

	database.DB.Preload("OrderItem.Product").First(&order, order.ID)
	return order
}

func DeleteOrder(id uuid.UUID) error {
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
