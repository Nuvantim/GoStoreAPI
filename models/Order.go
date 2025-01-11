package models

import (
	"github.com/google/uuid"
	"time"
	"gorm.io/gorm"
)

type Order struct {
	ID          string 		`json:"id" gorm:"type:uuid;primaryKey"`
	UserID      uint      `json:"user_id"`
	Total_Price uint      `json:"total_price"`
	Total_Item  uint      `json:"total_item"`
	Status      string    `json:"status" gorm:"type:enum('pending', 'paid', 'shipped', 'completed', 'canceled');default:'pending'"`
	OrderItem   []OrderItem
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

type OrderItem struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID    string    `json:"order_id" gorm:"type:uuid"`
	ProductID  uint      `json:"product_id"`
	Product    Product   `gorm:"foreignKey:ProductID"`
	Quantity   uint      `json:"quantity"`
	Total_Cost uint      `json:"total_cost"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error){
	o.ID = uuid.New().String()
	return
}
