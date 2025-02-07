package server

import (
	"api/internal/config"
	"api/internal/database"
	"api/internal/database/seeder"
	"api/internal/routes"
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func gracefulShutdown(app *fiber.App, done chan bool) {
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

func SetupAPI() {
	database.Setup()

	seeder.SeederSetup()

	app := fiber.New(config.FiberConfig())

	config.MiddlewareConfig(app)

	routes.Setup(app)
	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// app.Listen(":"+os.Getenv("PORT"), fiber.ListenConfig{EnablePrefork: true})
	app.Listen(":" + os.Getenv("PORT"))

	gracefulShutdown(app, done)

	<-done

	log.Println("Graceful shutdown complete.")

}
