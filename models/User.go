package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"PrimaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null" validate:"required"`
	Email     string    `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password  string    `json:"password" gorm:"not null" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type UserInfo struct {
	ID        uint      `json:"id" gorm:"PrimaryKey;autoIncrement"`
	UserID    uint      `json:"user_id"`
	Age       uint      `json:"age"`
	Phone     uint      `json:"phone"`
	District  string    `json:"district"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
