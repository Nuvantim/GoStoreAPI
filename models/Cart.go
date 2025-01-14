package models

import (
	"time"
)

type Cart struct {
	ID         uint      `json:"id" gorm:"PrimaryKey;autoIncrement"`
	UserID     uint      `json:"user_id" validate:"required"`
	ProductID  uint      `json:"product_id" validate:"required"`
	Product    Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity   uint      `json:"quantity" gorm:"default:1" validate:"required"`
	Total_Cost uint      `json:"total_cost"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
