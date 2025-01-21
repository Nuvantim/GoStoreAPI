package main

import (
	"api/config"
	"api/database"
	"api/database/seeder"
	"api/routes"
	"github.com/gofiber/fiber/v3"
	"os"
)

func main() {
	database.Setup()

	seeder.SeederSetup()

	app := fiber.New(config.FiberConfig())

	config.MiddlewareConfig(app)

	routes.Setup(app)

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	// app.Listen(":"+os.Getenv("PORT"), fiber.ListenConfig{EnablePrefork: true})
	app.Listen(":" + os.Getenv("PORT"))

}
