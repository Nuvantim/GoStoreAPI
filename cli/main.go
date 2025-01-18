package main

import (
	"api/config"
	"api/database"
	"api/routes"
	"github.com/gofiber/fiber/v3"
	"log"
	"os"
)

func main() {
	database.Setup()

	app := fiber.New(config.FiberConfig())

	config.MiddlewareConfig(app)

	routes.Setup(app)

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":"+os.Getenv("PORT")))

}
