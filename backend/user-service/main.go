package main

import (
    "log"
    "net/http"
    "os"
    "user-service/models" // Import your models package
    "user-service/routes" // Import user service routes
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

    // Set up routes for user service
    router := routes.SetupRoutes()

    // Set up CORS options for frontend origin
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:5173"}, // Replace with your frontend origin for production
        AllowCredentials: true,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
    })

    // Use the CORS middleware for the router
    handler := c.Handler(router)

    // Start the server with CORS-enabled router
    log.Println("User service is running on :8081...")
    if err := http.ListenAndServe(":8081", handler); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}
