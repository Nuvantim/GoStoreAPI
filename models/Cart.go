package models

import (
	"time"
)

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User       `gorm:"foreignKey:UserID"`
	Items     []CartItem `gorm:"foreignKey:ChartID"`
	CreatedAt time.Time
	UpdateAt  time.Time
}

type CartItem struct {
	ID        uint `gorm:primaryKey`
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID"`
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint
	CreatedAt time.Time
	UpdateAt  time.Time
}
