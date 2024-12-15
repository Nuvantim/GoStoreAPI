package models

import (
	"time"
)

type Cart struct {
	ID     uint  `json:"id" gorm:"PrimaryKey;autoIncrement"`
	UserID uint `json:"user_id"`
	// User       User    `gorm:"foreignKey:UserID"`
	ProductID  uint    `json:"product_id"`
	Product    Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity   uint    `json:"quantity" gorm:"default:1"`
	Total_Cost uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
