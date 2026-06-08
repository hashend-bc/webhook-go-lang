package handlers

import (
	"encoding/json"
	"net/http"
)

var events []map[string]interface{}

func Webhook(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)

	events = append(events, data)

	w.Write([]byte("Webhook received"))
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}