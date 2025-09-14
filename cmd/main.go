package main

import (
	"log"
	"os"

	"github.com/IbadT/tutor_app_back.git/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create and start server
	server, err := app.NewServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	// Start server
	log.Fatal(server.Start(port))
}
