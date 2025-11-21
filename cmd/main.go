package main

import (
	"api/internal/config"
	"api/internal/server/http"
	"api/pkg/guard"
	rds "api/redis"

	"log"
	"os"
)

func main() {
	// generate RSA
	guard.GenRSA()
	// Check RSA
	guard.CheckRSA()
	app := server.SetupAPI()

	done := make(chan bool, 1)

	go func() {
		app.Listen(":" + os.Getenv("PORT"))
	}()

	config.GracefulShutdown(app, done)

	<-done

	log.Println("Graceful shutdown complete.")

	// close redis
	rds.RedisClose()

}
