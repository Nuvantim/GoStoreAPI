package models

type User struct {
	ID       uint   `json:"id" gorm:"PrimaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Address  string `json:"address" gorm:"type:varchar(100);not null"`
	Phone    uint   `json:"phone" gorm:"not null"`
}
