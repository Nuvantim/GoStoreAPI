package models

import (
	"time"
)

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name      string    `json:"name" gorm:"not null" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
