package main

import (
	"log"
	"net/http"

	"github.com/xclamation/go-real-time-messenger/config"
	"github.com/xclamation/go-real-time-messenger/internal/routes"
)

func main() {
	// Load the cofiguration
	cfg := config.LoadConfig()
	
	// Log the environment we are in
	log.Printf("Starting server in %s mode\n", cfg.AppEnv())

	// Initialize the routes
	mux := routes.InitializeRoutes()

	// Start the server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
