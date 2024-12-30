package service

import (
	"api/database"
	"api/models"
)

func CreateReview(review models.Review) models.Product {
	var product models.Product
	database.DB.Create(&review)
	database.DB.Preload("Review").First(&product, review.ProductID)
	return product
}

func FindReview(ID uint) models.Review {
	var review models.Review
	database.DB.First(&review, ID)
	return review
}

func DeleteReview(ID uint) error {
	var review models.Review
	if err := database.DB.First(&review, ID).Error; err != nil {
		return err
	}
	database.DB.Delete(&review)
	return nil
}
