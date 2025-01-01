package models

import(
	
	"time"
)

type User struct {
	ID       	uint   		`json:"id" gorm:"PrimaryKey;autoIncrement"`
	Name     	string 		`json:"name" gorm:"not null"`
	Email    	string 		`json:"email" gorm:"not null"`
	Password 	string 		`json:"password" gorm:"not null"`
	CreatedAt       time.Time 	`gorm:"autoCreateTime"`
}

type UserInfo struct {
	ID              uint            `json:"id" gorm:"PrimaryKey;autoIncrement"`
	UserID          uint            `json:"user_id"`
	Age             uint            `json:"age"`
	Phone    	uint   		`json:"phone"`
	District        string          `json:"district"`
	City            string          `json:"city"`
	State           string          `json:"state"`
	Country         string          `json:"country"`
	CreatedAt       time.Time       `gorm:"autoCreateTime"`
	  
}
