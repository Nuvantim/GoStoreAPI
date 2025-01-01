package models

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name string `json:"name" gorm:"not null"`
}
