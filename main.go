package main

import (
	"fmt"
	"log"
	"net/http"
	"webhook-project/handlers"

)


func main() {
	http.HandleFunc("/webhook", handlers.WebhookHandler)
	http.HandleFunc("/events", handlers.GetEventsHandler)
	
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}