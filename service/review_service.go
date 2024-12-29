package service

import(
  "api/database"
  "api/models"
)

func CreateReview(review models.Review) model.Product{
  var product models.Product
  database.DB.Create(&review)
  database.DB.First(&product, review.ProductID).Preload("Review")
  return product
}

func FindReview(id uint) models.Review{
  var review models.Review
  database.DB.First(&review, id)
  return review
}

func DeleteReview(id uint) error {
  var review models.Review
  if err := database.DB.First(&review, id).Error; err != nil{
    return err
  }
  database.DB.Delete(&review)
}
