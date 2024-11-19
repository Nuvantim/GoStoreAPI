package models

type Category struct {
	ID      uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name    string    `json:"name" gorm:"not null"`
	Product []Product `json:"product" gorm:"foreignKey:CategoryID"`
}
