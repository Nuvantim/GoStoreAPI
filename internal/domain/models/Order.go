package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID         uuid.UUID   `json:"id" gorm:"primaryKey;type:char(36);not null;unique"`
	UserID     uint64      `json:"-" gorm:"not null"`
	TotalPrice uint64      `json:"total_price" gorm:"not null"`
	TotalItem  uint64      `json:"total_item" gorm:"not null"`
	Status     string      `json:"status" gorm:"default:'pending'"`
	OrderItems []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:OrderID"`
	CreatedAt  *time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

type OrderItem struct {
	ID        uint64    `json:"-" gorm:"primaryKey;autoIncrement"`
	OrderID   uuid.UUID `json:"order_id" gorm:"not null;type:char(36);references:ID"`
	ProductID uint64    `json:"product_id" gorm:"not null"`
	Product   Product   `json:"product,omitempty"`
	Quantity  uint64    `json:"quantity" gorm:"not null"`
	TotalCost uint64    `json:"total_cost" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.ID, _ = uuid.NewV7()
	return nil
}
