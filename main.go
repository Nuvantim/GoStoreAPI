package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
	"os"
	"toy-store-api/config"
	"toy-store-api/database"
	"toy-store-api/routes"
)

func main() {
	database.MysqlConnect()

	app := fiber.New(config.FiberConfig())

	config.MiddlewareConfig(app)

	routes.Setup(app)

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(os.Getenv("PORT")))

}
