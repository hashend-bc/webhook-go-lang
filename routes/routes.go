package routes

import (
	"net/http"
	"webhook-project/handlers"
	"webhook-project/middleware"
)

func SetupRoutes() {

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)

	http.HandleFunc("/add-ip", middleware.JWTMiddleware(handlers.AddIP))

	http.HandleFunc("/protected",
		middleware.JWTMiddleware(
			middleware.IPWhitelistMiddleware(handlers.Protected),
		),
	)

	http.HandleFunc("/webhook", handlers.Webhook)
	http.HandleFunc("/events", handlers.GetEvents)
}