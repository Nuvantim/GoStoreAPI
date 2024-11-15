package models

import (
	"time"
)

type Riview struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User `gorm:"foreignKey;UserID"`
	ProductID uint
	Product   Product `gorm:"foreignKey;ProductID"`
	Rating    int     `gorm:"check:rating >= 1 AND rating <= 5"`
	Comment   string
	CreatedAt time.Time
}
