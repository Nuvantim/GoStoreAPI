package service

import (
	"api/database"
	"api/models"
	"errors"
	"gorm.io/gorm"
)

type Cart = models.Cart 		// declare type models Cart

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

func FindCart(id uint) Cart {
	var cart Cart         		// declare variabel Cart
	database.DB.Where("id = ?", id).
		Preload("Product", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "price", "stock", "category_id")
		}).Preload("Product.Category").Take(&cart)
	return cart
}

func CreateCart(cart_data Cart, id_user, cost uint) Cart {
	cart := Cart{
		UserID:     id_user,
		ProductID:  cart_data.ProductID,
		Total_Cost: cost,
	}
	database.DB.Create(&cart)
	carts := FindCart(cart.ID)
	return carts
}

func UpdateCart(cart_update Cart, cost uint) Cart {
	var cart Cart         		// declare variabel Cart
	//Get cart by ID
	database.DB.Take(&cart, cart_update.ID)

	// Update data
	cart.Quantity = cart_update.Quantity
	cart.Total_Cost = cost
	database.DB.Save(&cart)

	// return data
	carts := FindCart(cart.ID)
	return carts
}

func DeleteCart(input interface{}) error {
	var cart Cart         		// declare variabel Cart
	switch v := input.(type) {
	case uint:
		return database.DB.Where("id = ?", v).Delete(&cart).Error

	case []uint:
		return database.DB.Where("id IN ?", v).Delete(&cart).Error

	default:
		return errors.New("invalid input type")
	}
}

func TransferCart(cart_id []uint) []Cart {
	var carts []Cart
	result := database.DB.Where("id IN ?", cart_id).Limit(1).Find(&carts)
	if result.RowsAffected == 0 {
		return nil // return nil if data empty
	}

	database.DB.Where("id IN ?", cart_id).Find(&carts)
	return carts
}
