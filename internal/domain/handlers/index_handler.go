package handler

import (
	"github.com/gofiber/fiber/v2"
  _"os"
)

func Home(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{})
}

func Swagger(c *fiber.Ctx) error {
	// var url string
	// url = os.Getenv("APP_URL")
	// if url == "" {
	//   url = "127.0.0.1"+ ":" + os.Getenv("PORT")
	// }
	// url = "127.0.0.1" + ":" + os.Getenv("PORT")
	// return c.Render("api", fiber.Map{
	// 	"Title":      "GoStoreAPI Documentation",
	// 	"APIBaseURL": url,
	// })
  return c.Render("api", fiber.Map{})
}
