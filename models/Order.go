package models

import (
	// "github.com/google/uuid"
	// "gorm.io/gorm"
	"time"
)

type Order struct {
	ID         uint        `json:"id" gorm:"autoIncrement;primaryKey"`
	UserID     uint        `json:"user_id" gorm:"not null"`
	TotalPrice uint        `json:"total_price" gorm:"not null"`
	TotalItem  uint        `json:"total_item" gorm:"not null"`
	Status     string      `json:"status" gorm:"type:enum('pending','paid','shipped','completed','canceled');default:'pending'"`
	OrderItems []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID   uint      `json:"order_id" gorm:"type:uuid;not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Product   *Product  `json:"product,omitempty"`
	Quantity  uint      `json:"quantity" gorm:"not null"`
	TotalCost uint      `json:"total_cost" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// func (o *Order) BeforeCreate(tx *gorm.DB) error {
// 	if o.ID == uuid.Nil {
// 		o.ID = uuid.New()
// 	}
// 	return nil
// }
