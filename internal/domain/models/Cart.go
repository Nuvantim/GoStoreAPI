package models

import (
	"time"
)

type Cart struct {
	ID         uint64    `json:"id" gorm:"PrimaryKey;autoIncrement"`
	UserID     uint64    `json:"-"`
	ProductID  uint64    `json:"product_id"`
	Product    Product   `json:"product" gorm:"foreignKey:ProductID"`
	Quantity   uint64    `json:"quantity" gorm:"default:1"`
	Total_Cost uint64    `json:"total_cost"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
