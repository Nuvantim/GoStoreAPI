package config

import (
	"context"
	"log"
	"os"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func GracefulShutdown(app *fiber.App, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

// FiberConfig berisi konfigurasi Fiber yang aman
func FiberConfig() fiber.Config {
	engine := html.New("views", ".html")
	return fiber.Config{
		AppName:       "GoStoreAPI",
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Nuvantim Project",
		Prefork:       true,
		Views:         engine,
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

	//Helmet
	app.Use(helmet.New(helmet.Config{
		ContentSecurityPolicy: "dafault-src 'self'; frame-ancestors 'self'",
		HSTSMaxAge:            31536000,
		HSTSPreloadEnabled:    true,
		HSTSExcludeSubdomains: false,
	}))

	//Idempotency
	app.Use(idempotency.New())

	// CSRF Protection
	// app.Use(csrf.New())

	// CORS Configuration
	var url string = os.Getenv("URL")
	var port string = os.Getenv("PORT")

	var origin = fmt.Sprintf("%s,http://localhost:%s, http://127.0.0.1:%s",url,port,port)
	app.Use(cors.New(cors.Config{
		AllowOrigins: origin,
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Authorization, Accept",
		MaxAge:       3600,
	}))
}
