package main

import (
	"api/internal/config"
	"api/internal/database"
	rds "api/internal/redis"
	"api/internal/server/http"
	"api/pkg/guard"

	"log"
	"os"
)

func main() {
	// generate RSA
	guard.GenRSA()
	// Check RSA
	guard.CheckRSA()
	app := server.SetupAPI()

	done := make(chan struct{}, 1)

	go func() {
		if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
			log.Printf("server error: %v", err)
			select {
			case done <- struct{}{}:
			default:
			}
		}
	}()

	go config.GracefulShutdown(app, done)

	<-done

	log.Println("Graceful shutdown complete.")

	// close redis
	rds.RedisClose()

	// close database
	database.Close()

}
