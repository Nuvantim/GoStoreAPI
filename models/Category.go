package models

type Category struct {
	ID      uint      `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name    string    `json:"name" gorm:"type:varchar(100);not null"`
	Product []Product `json:"product" gorm:"foreignKey:CategoryID"`
}
