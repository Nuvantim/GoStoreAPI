package service

import (
	"api/database"
	"api/models"
)

func CreateReview(review_data models.Review) models.Product {
	review := models.Review{
		UserID:    review_data.UserID,
		ProductID: review_data.ProductID,
		Rating:    review_data.Rating,
		Comment:   review_data.Comment,
	}
	database.DB.Create(&review)

	var product models.Product
	database.DB.Preload("Review").Preload("User").First(&product, review.ProductID)

	return product
}

func FindReview(user_id uint) models.Review {
	var review models.Review
	database.DB.First(&review, user_id)
	return review
}

func FindUserReview(user_id, product_id uint) models.Review {
	var review models.Review
	database.DB.Where("user_id = ?", user_id).Where("product_id = ?", product_id).First(&review)
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
