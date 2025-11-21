package models

import (
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty" gorm:"PrimaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null" validate:"required"`
	Email     string    `json:"email,omitempty" gorm:"unique;not null" validate:"required,email"`
	Password  string    `json:"-" gorm:"not null"`
	Roles     []Role    `json:"roles" gorm:"many2many:user_roles;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type UserInfo struct {
	ID        uint64    `json:"-" gorm:"PrimaryKey;autoIncrement"`
	UserID    uint64    `json:"-"`
	Age       uint64    `json:"age"`
	Phone     uint64    `json:"phone"`
	District  string    `json:"district"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Country   string    `json:"country"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Token struct {
	Otp   uint64 `json:"otp"`
	Email string `json:"email"`
}
