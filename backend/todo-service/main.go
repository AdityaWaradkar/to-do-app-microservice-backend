package main

import (
	"log"
	"net/http"
	"os"
	"todo-service/models"
	"todo-service/routes"

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
		AllowedOrigins: []string{"http://to-do-list-web-app.s3-website.ap-south-1.amazonaws.com"}, // Allow all origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(router)

	log.Println("Todo service is running on :8082...")
	if err := http.ListenAndServe(":8082", handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
