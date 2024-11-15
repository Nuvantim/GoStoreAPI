package models

type Order struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User   User        `gorm:"foreignKey;UserID"`
	Total  float64     `gorm:"not null"`
	Status string      `gorm:"type:enum('pending','paid','shipped','completed','canceled';default:'pending"`
	Items  []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	Order     Order `gorm:"foreignKey:OrderID"`
	ProductID uint
	Product Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
}
