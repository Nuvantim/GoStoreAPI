package config

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"time"
	// "github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// FiberConfig berisi konfigurasi Fiber yang aman
func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:       "fiber-api",
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
	}
}

// MiddlewareSetup menyiapkan semua middleware keamanan
func MiddlewareConfig(app *fiber.App) {

	// Rate Limiting
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))
	// Logger
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
	// app.Use(logger.New())

	//Helmet
	app.Use(helmet.New())

	//Idempotency
	app.Use(idempotency.New())

	// CSRF Protection
	// app.Use(csrf.New())

	// CORS Configuration
	app.Use(cors.New())
}
