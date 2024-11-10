package models

import (
	"gorm.io/gorm"
)

type Category struct {
	ID      uint      `gorm:"PrimaryKey"`
	Name    string    `gorm:"size:100; not null"`
	Product []Product `gorm:"foreignKey:CategoryID"`
}
