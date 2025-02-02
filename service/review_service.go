package service

import (
	"api/database"
	"api/models"
)
type Review = models.Review  //declare type models Review

func CreateReview(review_data Review) Product {
	review := Review{
		UserID:    review_data.UserID,
		ProductID: review_data.ProductID,
		Rating:    review_data.Rating,
		Comment:   review_data.Comment,
	}
	database.DB.Create(&review)

	var product Product
	database.DB.Preload("Review").Preload("User").Take(&product, review.ProductID)

	return product
}

func FindReview(user_id uint) Review {
	var review Review
	database.DB.Take(&review, user_id)
	return review
}

func FindUserReview(user_id, product_id uint) Review {
	var review Review
	database.DB.Where("user_id = ?", user_id).Where("product_id = ?", product_id).Take(&review)
	return review
}

func DeleteReview(ID uint) error {
	var review Review
	if err := database.DB.Take(&review, ID).Error; err != nil {
		return err
	}
	database.DB.Delete(&review)
	return nil
}
