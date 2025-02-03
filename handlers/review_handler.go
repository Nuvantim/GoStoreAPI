package handler

import (
	"api/service"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

type Review = service.Review //declare type models Review

/*
HANDLER Create Review
*/
func CreateReview(c fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint)

	//parse body to json
	var review Review
	if err := c.Bind().Body(&review); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// check user review
	user_review := service.FindUserReview(user_id, review.ProductID)
	if user_review.ID != 0 {
		return c.Status(403).JSON(fiber.Map{
			"message": "User Review Already Exist",
		})
	}

	//check product
	product := service.FindProduct(review.ProductID)
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product Not Found",
		})
	}

	//attach user_id & product_id
	review.UserID = user_id

	//connect to service
	reviews := service.CreateReview(review)
	return c.Status(200).JSON(reviews)
}

/*
HANDLER DELETE REVIEW
*/
func DeleteReview(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user_id := c.Locals("user_id").(uint)

	//check review
	review := service.FindReview(uint(id))
	switch {
	case review.ID == 0:
		return c.Status(404).JSON(fiber.Map{
			"message": "Review Not Found",
		})
	case review.UserID != user_id:
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	//connect to service
	service.DeleteReview(uint(id))
	return c.Status(200).JSON(fiber.Map{
		"message": "sucess",
	})
}
