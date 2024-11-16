package main

import (
	"log"
	"net/http"
	"os"
	"user-service/models"
	"user-service/routes"

	"github.com/rs/cors"
)


func main() {
	// Get the MongoDB URI from the environment variable
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	// Connect to MongoDB
	err := models.ConnectDB(mongoURI)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %s\n", err)
	}

	// Set up routes
	router := routes.SetupRoutes()

	// Set up CORS with credentials support and allow origin from your frontend
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://to-do-list-app-9876.s3-website-us-east-1.amazonaws.com", // Updated frontend URL
		}, // Frontend URL
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
		AllowCredentials: true, // Allow credentials (cookies, sessions)
	})

	// Apply CORS middleware
	handler := corsHandler.Handler(router)

	log.Println("User service is running on :8081...")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
