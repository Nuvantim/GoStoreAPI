package main

import (
	"log"

	"e-commerce-api/config"
	"e-commerce-api/database"
	"e-commerce-api/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	database.MysqlConnect()

	app := fiber.New(config.FiberConfig())

	config.MiddlewareSetup(app)

	routes.Setup(app)

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))

}
