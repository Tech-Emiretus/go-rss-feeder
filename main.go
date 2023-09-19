package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/techemeritus/rss-feeder/router"
)

func main() {
	log.Println("RSS Feeder Application Starting...")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not read environment variables from .env file. Error: ", err)
	}

	appPort := os.Getenv("APP_PORT")

	if appPort == "" {
		log.Fatal("Application Port cannot be empty.")
	}

	log.Println("Listening on Port: ", appPort)
	err := http.ListenAndServe(":" + appPort, router.GetRouter())

	if err != nil {
		log.Fatal("RSS Feeder Application Failed to start. Error: ", err)
	}
}