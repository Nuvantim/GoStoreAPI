package config

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
        "github.com/go-playground/validator/v10"
	"time"
	// "github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
)
//implement the Validate method
func (v *structValidator) Validate(out any) error {
    return v.validate.Struct(out)
}

// FiberConfig berisi konfigurasi Fiber yang aman
func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:       "fiber-api",
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		StructValidator: &structValidator{validate: validator.New()},
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
	app.Use(logger.New())

	//Helmet
	app.Use(helmet.New())

	//Idempotency
	app.Use(idempotency.New())

	// CSRF Protection
	// app.Use(csrf.New())

	// CORS Configuration
	app.Use(cors.New())
}
