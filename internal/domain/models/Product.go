package models

import (
	"time"
)

type Product struct {
	ID          uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null" validate:"required"`
	Description string    `json:"description" gorm:"not null" validate:"required"`
	Price       uint64    `json:"price" gorm:"not null" validate:"required"`
	Stock       uint64    `json:"stock" gorm:"default:0"`
	CategoryID  uint64    `json:"category_id" gorm:"not null" validate:"required"`
	Category    *Category `json:"category" gorm:"foreignKey:CategoryID"`
	Review      []Review  `json:"review"`
	CreatedAt   time.Time `json:"CreatedAt" gorm:"autoCreateTime"`
}
