package main

import (
	"gateway/middleware"
	"gateway/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	loggedMux := middleware.LoggingMiddleware(mux)

	log.Println("🚀 Gateway running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", loggedMux))
}
