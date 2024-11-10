package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"PrimaryKey"`
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100;unique;not null"`
	Password string `gorm:"not null"`
	Address  string
	Phone    string `gorm:"size:15"`
}
