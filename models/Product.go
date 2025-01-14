package models

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id" gorm:"PrimaryKey;autoIncrement"`
	Name        string    `json:"name" gorm:"not null" validate:"required"`
	Description string    `json:"description" gorm:"not null" validate:"required"`
	Price       uint      `json:"price" gorm:"not null" validate:"required"`
	Stock       uint      `json:"stock" gorm:"default:0"`
	CategoryID  uint      `json:"category_id" gorm:"not null" validate:"required"`
	Category    Category  `json:"category, omitempty" gorm:"foreignKey:CategoryID"`
	Review      *[]Review `json:"review,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
