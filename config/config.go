package config

import (
	"log"
	"time"
        "os"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	// "github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/helmet"
)

// FiberConfig berisi konfigurasi Fiber yang aman
func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:       "fiber-api",
		CaseSensitive: true,
		StrictRouting: true,
		EnablePrefork : true,
		ServerHeader:  "Fiber",
	}
}

// MiddlewareSetup menyiapkan semua middleware keamanan
func MiddlewareConfig(app *fiber.App) {
	// Recover dari panic
	app.Use(recover.New())

	// Logging
	app.Use(requestid.New())
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	defer logFile.Close()
	app.Use(logger.New(logger.Config{
			Format: "${locals:requestid} ${ip} ${status} - ${method} ${path} ${body}\n",
			Output: logFile,
	}))

	// Rate Limiting
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))
	
	//Helmet
	app.Use(helmet.New())
	
	//Idempotency
	app.Use(idempotency.New())

	// CSRF Protection
	// app.Use(csrf.New())

	// CORS Configuration
	app.Use(cors.New())
}
