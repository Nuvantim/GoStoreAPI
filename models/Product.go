package models

type Product struct {
	ID          uint     `json:"id" gorm:"PrimaryKey"`
	Name        string   `json:"name" gorm:"not null"`
	Description string   `json:"description" gorm:"not null"`
	Price       float64  `json:"price" gorm:"not null"`
	Stock       int      `json:"stock" gorm:"default:0"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}
