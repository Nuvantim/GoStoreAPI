package service

import (
	"api/internal/database"
	"api/internal/domain/models"
	"errors"
	"gorm.io/gorm"
)

type Cart = models.Cart // declare type models Cart

func GetCart(id uint) []Cart {
	var cart []Cart
	database.DB.Where("user_id = ?", id).
		Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "price", "stock", "category_id")
		}).
		Preload("Product.Category").
		Find(&cart)
	return cart
}


func FindCart(id interface{}) (Cart, []Cart) {
	var ( // declare variabel Cart
		cart  Cart
		carts []Cart
	)
	switch CartID := id.(type) {

	case uint: //single id
		database.DB.Where("id = ?", CartID).
			Preload("Product", func(db *gorm.DB) *gorm.DB {
				return db.Select("id", "name", "price", "stock", "category_id")
			}).Preload("Product.Category").Take(&cart)

	case []uint: //multiple id
		database.DB.Where("id IN ?", CartID).Find(&carts)

	default:
		return cart, carts

	}
	return cart, carts

}

func CreateCart(cart_data Cart, id_user, cost uint) (Cart) {
	cart := Cart{
		UserID:     id_user,
		ProductID:  cart_data.ProductID,
		Total_Cost: cost,
	}
	database.DB.Create(&cart)
	// return cart data
	carts,_:= FindCart(cart.ID)
	return carts
}

func UpdateCart(cart_update Cart, cost uint) Cart {
	var cart Cart // declare variabel Cart
	//Get cart by ID
	database.DB.Take(&cart, cart_update.ID)

	// Update data
	cart.Quantity = cart_update.Quantity
	cart.Total_Cost = cost
	database.DB.Save(&cart)

	// return cart data
	carts,_:= FindCart(cart.ID)
	return carts
}

func DeleteCart(input interface{}) error {
	var cart Cart // declare variabel Cart
	switch v := input.(type) {
	case uint:
		return database.DB.Where("id = ?", v).Delete(&cart).Error

	case []uint:
		return database.DB.Where("id IN ?", v).Delete(&cart).Error

	default:
		return errors.New("invalid input type")
	}
}
