package handler

import(
  "api/service"
  "api/models"
  "github.com/gofiber/fiber/v3"
  "strconv"
)

/*
HANDLER Create Review
*/
func CreateReview(c fiber.Ctx) error{
  user_id := c.Locals(user_id).(uint)
  product_id := strconv.Atoi(c.Params("id"))

  //check product
  product := service.FindProduct(uint(product_id))
  if product.ID == 0 {
    return c.Status(404).JSON(fiber.Map{
      "message" : "Product Not Found",
    })
  }
  //parse body to json
  var review models.Review
  if err := c.Bind().Body(review); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
  }

//connect to service
reviews := service.CreateReview(user_id,uint(product_id),review)
return c.Status(200).JSON(reviews)
}

/*
HANDLER DELETE REVIEW
*/
func DeleteReview(c fiber.Ctx) error {
  id := strconv.Atoi(c.Params("id"))
  user_id := c.Locals("user_id").(uint)

  //check review
  review := service.FindReview("id")
  if review.UserId != user_id {
    return c.status(403).JSON(fiber.Map{
      "message" : "Forbidden",
    })
  }
  
  //connect to service
  service.DeleteReview(uint(id))
  return c.Status(200).JSON(fiber.Map{
    "message":"sucess",
  }
}
