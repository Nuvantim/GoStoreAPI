package service

import (
	"api/database"
	"api/models"
)

func CreateReview(review models.Review) models.Product {
	database.DB.Create(&review)

	var product models.Product
	database.DB.Preload("Review").Preload("User").First(&product, review.ProductID)

	return product
}

func FindReview(ID uint) models.Review {
	var review models.Review
	database.DB.First(&review, ID)
	return review
}

func FindUserReview(ID uint) models.Review {
	var review models.Review
	database.DB.Where("user_id = ?", ID).First(&review)
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
