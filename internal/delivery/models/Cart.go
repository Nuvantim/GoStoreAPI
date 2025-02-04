package models

import (
	"time"
)

type Cart struct {
	ID         uint      `json:"id" gorm:"PrimaryKey;autoIncrement"`
	UserID     uint      `json:"-"`
	ProductID  uint      `json:"product_id"`
	Product    Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity   uint      `json:"quantity" gorm:"default:1"`
	Total_Cost uint      `json:"total_cost"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
