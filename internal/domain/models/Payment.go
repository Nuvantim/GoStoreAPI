package models

import (
	"time"
)

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	OrderID       uint      `gorm:"not null"`
	Order         Order     `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount        float64   `gorm:"not null"`
	PaymentDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	PaymentMethod string    `gorm:"type:varchar(20);not null;check:payment_method IN ('bank_transfer', 'qris')"`
	Status        string    `gorm:"type:varchar(20);default:'pending';check:status IN ('pending', 'completed', 'failed')"`
}
