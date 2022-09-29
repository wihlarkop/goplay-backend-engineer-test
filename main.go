package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"goplay-backend-engineer-test/container"
	"goplay-backend-engineer-test/handlers"
	"goplay-backend-engineer-test/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	srv, err := server.NewGinHttpRouter(
		&server.Config{
			Address: os.Getenv("PORT"),
		})

	if err != nil {
		panic(err)
	}

	containers := container.NewContainer()

	router := handlers.NewRouter(srv.Router, containers)
	router.RegisterRouter()

	srv.Start()
}
