package main

import (
	"log"
	"net/http"

	"webhook-project/config"
	"webhook-project/routes"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	config.ConnectDB()
	routes.SetupRoutes()

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}