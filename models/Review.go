package models

import (
	"time"
)

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json: "user_id" gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Rating    int       `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment   string    `json:"comment" gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
