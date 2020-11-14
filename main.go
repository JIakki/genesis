package main

import (
	"context"
	"github.com/JIakki/genesis/api"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Please provide PORT variable")
	}

	srv := api.Create(&api.Params{Port: port})

	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
