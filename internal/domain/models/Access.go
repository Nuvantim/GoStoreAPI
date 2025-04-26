package models

import "time"

type Role struct {
	ID          uint64       `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"unique" validate:"required"`
	CreatedAt   time.Time    `json:"created_at"`
	Permissions []Permission `json:"permissions,omitempty" gorm:"many2many:role_permissions;" validate:"required"`
}

type Permission struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
