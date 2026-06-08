package handlers

import (
	"encoding/json"
	"net/http"
	"webhook-project/models"
	"webhook-project/services"
)

func AddIP(w http.ResponseWriter, r *http.Request) {
	var ip models.WhitelistIP
	json.NewDecoder(r.Body).Decode(&ip)

	ip.Enabled = true
	services.AddIP(ip)

	w.Write([]byte("IP added"))
}