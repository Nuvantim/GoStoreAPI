package service

import (
	"api/internal/database"
	"api/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ( //declare type models Order & OrderItems
	Order     = models.Order
	OrderItem = models.OrderItem
)

/*
SERVICE ORDER
*/
func GetOrder(id uint64) []Order {
	var order []Order
	database.DB.Where("user_id = ?", id).Find(&order)
	return order
}

func FindOrder(id uuid.UUID) Order {
	var order Order
	database.DB.Preload("OrderItems.Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "price", "stock")
	}).Take(&order, "id = ?", id)
	return order
}

func CreateOrder(id_user, totalItem, totalPrice uint64, cartData []Cart) Order {
	// Create Order
	order := Order{
		UserID:     id_user,
		TotalPrice: totalPrice,
		TotalItem:  totalItem,
	}
	database.DB.Create(&order)

	// Create Order Item
	var orderItems []OrderItem

	// Menyiapkan data untuk batch insert
	for _, cart := range cartData {
		orderItems = append(orderItems, OrderItem{
			OrderID:   order.ID,
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			TotalCost: cart.Total_Cost,
		})
	}

	// Batch insert ke database
	database.DB.Create(&orderItems)

	database.DB.Take(&order, order.ID)
	return order
}

func DeleteOrder(id uuid.UUID) error {
	// start transaction database
	tx := database.DB.Begin()

	// Hapus OrderItem berdasarkan OrderID
	if err := tx.Where("order_id = ?", id).Delete(&OrderItem{}).Error; err != nil {
		tx.Rollback() // cancel transaction if error
		return err
	}

	// Delete Order by ID
	if err := tx.Delete(&Order{}, id).Error; err != nil {
		tx.Rollback() // cancel transaction if error
		return err
	}

	// Commit transaksi jika semua berhasil
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
