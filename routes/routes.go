package routes

import (
	"github.com/gofiber/fiber/v3"
	"toy-store-api/handlers"
)

func Setup(app *fiber.App) {

	//auth Route
	app.Post("/login", handler.Login)
	app.Post("/logout", handler.Logout)

	// User Route
	app.Get("/user", handler.GetUser)
	app.Post("/register", handler.RegisterUser)
	app.Get("/user/:id", handler.FindUser)
	app.Put("/user/update/:id", handler.UpdateUser)
	app.Delete("/user/delete/:id", handler.DeleteUser)

	// Category Route
	app.Get("/category", handler.GetCategory)
	app.Get("/category/:id", handler.FindCategory)
	app.Post("/category/store", handler.CreateCategory)
	app.Put("/category/update/:id", handler.UpdateCategory)
	app.Delete("/category/delete/:id", handler.DeleteCategory)

	// Product Route
	app.Get("/product", handler.GetProduct)
	app.Get("/product/:id", handler.FindProduct)
	app.Post("/product/store", handler.CreateProduct)
	app.Put("/product/update/:id", handler.UpdateProduct)
	app.Delete("/product/delete/:id", handler.DeleteProduct)

	// app.Put("/book/:id", handler.UpdateBooks)
	// app.Delete("/book/:id", handler.DeleteBooks)
}
