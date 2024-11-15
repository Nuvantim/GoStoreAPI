package routes

import(
	"github.com/gofiber/fiber/v3"
	"e-commerce-api/handlers"
)

func Setup(app *fiber.App){

	app.Get("/category", handler.GetCategory)
	app.Get("/category/:id", handler.FindCategory)
	app.Post("/category/store", handler.CreateCategory)

	// app.Put("/book/:id", handler.UpdateBooks)
	// app.Delete("/book/:id", handler.DeleteBooks)
}
