package main

import (
	"api/internal/server/http"
	"api/internal/config"
	"api/pkg/guard"

	"os"
	"log"
)

func main() {
	guard.GenRSA()
	app := server.SetupAPI()

	done := make(chan bool, 1)

	go func() {
		app.Listen(":" + os.Getenv("PORT"))
	}()

	config.GracefulShutdown(app, done)

	<-done

	log.Println("Graceful shutdown complete.")
}
