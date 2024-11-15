package models

import (
	"time"
)

type Payment struct {
	ID            uint `gorm:"primaryKey"`
	OrderID       uint
	Order         Order     `gorm:"foreignKey:OrderID"`
	Amount        float64   `gorm:"not null"`
	PaymentDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	PaymentMethod string    `gorm:"type:enum('bank_transfer','qris');not null"`
	Status        string    `gorm:"type:enum('pending','complated','failed');default:'pending'"`
}
