package server

import (
	"api/internal/config"
	"api/internal/database"
	"api/internal/database/seeder"
	rds "api/internal/redis"
	"api/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupAPI() *fiber.App {
	// start database setup
	database.Setup()

	// run database seeder
	seeder.SeederSetup()

	// run redis
	rds.InitRedis()

	// start fiber configuration
	app := fiber.New(config.FiberConfig())

	// start middlewar config
	config.MiddlewareConfig(app)

	// start route setup
	routes.Setup(app)

	return app
}
