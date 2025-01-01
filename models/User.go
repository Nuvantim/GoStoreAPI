package models

import(
	
	"time"
)

type User struct {
	ID       	uint   		`json:"id" gorm:"PrimaryKey;autoIncrement"`
	Name     	string 		`json:"name" gorm:"not null"`
	Email    	string 		`json:"email" gorm:"not null"`
	Password 	string 		`json:"password" gorm:"not null"`
	Address  	string 		`json:"address" gorm:"not null"`
	Phone    	uint   		`json:"phone" gorm:"not null"`
	CreatedAt   time.Time 	`gorm:"type:date"`
}
