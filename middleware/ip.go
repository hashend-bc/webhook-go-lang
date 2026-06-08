package middleware

import (
	"context"
	"net"
	"net/http"
	"time"

	"webhook-project/config"
	"webhook-project/models"
)

func GetClientIP(r *http.Request) string {
	ip := r.RemoteAddr
	host, _, _ := net.SplitHostPort(ip)
	return host
}

func IPWhitelistMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ip := GetClientIP(r)

		var result models.WhitelistIP

		err := config.DB.Collection("whitelist").
			FindOne(context.TODO(), map[string]interface{}{
				"ip_address": ip,
				"enabled":    true,
			}).Decode(&result)

		status := "SUCCESS"
		errMsg := ""

		if err != nil {
			status = "DENIED"
			errMsg = "IP not allowed"
			http.Error(w, errMsg, http.StatusForbidden)
		} else {
			next(w, r)
		}

		log := models.AccessLog{
			IPAddress: ip,
			Endpoint:  r.URL.Path,
			Status:    status,
			Error:     errMsg,
			Time:      time.Now(),
		}

		config.DB.Collection("logs").InsertOne(context.TODO(), log)
	}
}