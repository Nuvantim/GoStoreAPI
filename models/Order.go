package models

type Order struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint   `json:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Total  uint   `json:"total"`
	Status string `json:"status" gorm:"type:enum('pending', 'paid', 'shipped', 'completed', 'canceled');default:'pending'"`
	
}

type OrderItem struct {
	ID         uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID    uint    `json:"order_id"`
	Order      Order   `json:"order" gorm:"foreignKey:OrderID"`
	ProductID  uint    `json:"product_id"`
	Product    Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity   uint    `json:"quantity"`
	Total_Cost uint    `json:"total_cost"`
}
