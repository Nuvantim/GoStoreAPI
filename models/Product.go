package models


type Product struct {
	ID          uint    `gorm:"PrimaryKey"`
	Name        string  `gorm:"type:size(100);not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"type:not null"`
	Stock       int     `gorm:"default:0"`
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
}
