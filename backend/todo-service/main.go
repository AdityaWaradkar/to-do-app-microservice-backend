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
    // Load MongoDB URI from environment variable
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI environment variable not set")
    }

    // Connect to MongoDB
    if err := models.ConnectDB(mongoURI); err != nil {
        log.Fatalf("Could not connect to MongoDB: %s\n", err)
    }
    log.Println("Connected to MongoDB successfully")

    // Set up the routes
    router := routes.SetupRoutes()

    // Configure CORS
    corsHandler := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:5173"},  // React app's URL
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type"},
        AllowCredentials: true,
    })

    // Apply CORS handler
    handler := corsHandler.Handler(router)

    // Start the HTTP server
    log.Println("Todo service is running on :8082...")
    if err := http.ListenAndServe(":8082", handler); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}
