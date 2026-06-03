package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"webhook-project/models"
	"webhook-project/storage"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	storage.Events = append(storage.Events, event)

	fmt.Println("Received event:", event.Type)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Event received"))
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(storage.Events); err != nil {
		fmt.Println("Error encoding events:", err)
	}
}
