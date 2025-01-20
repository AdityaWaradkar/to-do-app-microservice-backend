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
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	err := models.ConnectDB(mongoURI)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %s\n", err)
	}

	router := routes.SetupRoutes()

	// Update CORS configuration to allow all origins
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://to-do-list-web-app.s3-website.ap-south-1.amazonaws.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	log.Println("User service is running on :8081...")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
