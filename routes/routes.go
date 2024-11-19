package routes

import (
	"e-commerce-api/handlers"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {

	app.Get("/category", handler.GetCategory)
	app.Get("/category/:id", handler.FindCategory)
	app.Post("/category/store", handler.CreateCategory)
	app.Put("/category/update/:id", handler.UpdateCategory)
	app.Delete("/category/delete/:id", handler.DeleteCategory)

	app.Get("/product", handler.GetProduct)
	app.Get("/product/:id", handler.FindProduct)
	app.Post("/product/store", handler.CreateProduct)
	app.Put("/product/update/:id", handler.UpdateProduct)
	app.Delete("/product/delete/:id", handler.DeleteProduct)

	// app.Put("/book/:id", handler.UpdateBooks)
	// app.Delete("/book/:id", handler.DeleteBooks)
}
