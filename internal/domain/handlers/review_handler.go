package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils/responses"
	"api/pkg/utils/validates"
	"github.com/gofiber/fiber/v2"
)

type Review = service.Review //declare type models Review

/*
HANDLER Create Review
*/
func CreateReview(c *fiber.Ctx) error {
	user_id := c.Locals("user_id").(uint64)

	//parse body to json
	var review Review
	if err := c.BodyParser(&review); err != nil {
		return c.Status(400).JSON(response.Error("failed parser json", err.Error()))
	}

	// validate data
	if err := validate.BodyStructs(review); err != nil {
		return c.Status(422).JSON(response.Error("failed validate data", err.Error()))
	}

	// check user review
	user_review := service.FindUserReview(user_id, review.ProductID)
	if user_review.ID != 0 {
		return c.Status(403).JSON(response.Error("failed create review", "user review already exist"))
	}

	//check product
	product := service.FindProduct(review.ProductID)
	if product.ID == 0 {
		return c.Status(404).JSON(response.Error("failed find product", "product not found"))
	}

	//attach user_id & product_id
	review.UserID = user_id

	//connect to service
	reviews := service.CreateReview(review)
	return c.Status(200).JSON(response.Pass("success create review", reviews))
}

/*
HANDLER DELETE REVIEW
*/
func DeleteReview(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id < 0 {
		return c.Status(400).JSON(response.Error("failed get id", err.Error()))
	}
	user_id := c.Locals("user_id").(uint64)

	//check review
	review := service.FindReview(uint64(id))
	switch {
	case review.ID == 0:
		return c.Status(404).JSON(response.Error("failed find review", "review not found"))
	case review.UserID != user_id:
		return c.Status(403).JSON(response.Error("failed find review", "review forbidden"))
	}

	//connect to service
	if err := service.DeleteReview(uint64(id)); err != nil {
		return c.Status(500).JSON(response.Error("failed delete review", err.Error()))
	}
	return c.Status(200).JSON(response.Pass("review deleted", struct{}{}))
}
