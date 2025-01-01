package models

import (
	"time"
)

type Order struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID"`
	Total     uint   `json:"total"`
	Status    string `json:"status" gorm:"type:enum('pending', 'paid', 'shipped', 'completed', 'canceled');default:'pending'"`
	OrderItem []OrderItem
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type OrderItem struct {
	ID         uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID    uint    `json:"order_id" gorm:"not null"`
	ProductID  uint    `json:"product_id"`
	Product    Product `gorm:"foreignKey:ProductID"`
	Quantity   uint    `json:"quantity"`
	Total_Cost uint    `json:"total_cost"`
}
