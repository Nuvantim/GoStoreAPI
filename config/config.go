package config

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
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
func MiddlewareSetup(app *fiber.App) {
	// Recover dari panic
	app.Use(recover.New())

	// Logging
	app.Use(logger.New())

	// Rate Limiting
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))

	// CSRF Protection
	app.Use(csrf.New())

	// CORS Configuration
	app.Use(cors.New())
}
