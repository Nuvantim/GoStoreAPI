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
		if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
			log.Fatal("server error : ", err)
			done <- true
		}
	}()

	config.GracefulShutdown(app, done)

	<-done

	log.Println("Graceful shutdown complete.")

	// close redis
	rds.RedisClose()

}
