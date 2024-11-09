package main

import (
    "log"
    "net/http"
    "os"
    "todo-service/models" // Import the models package
    "todo-service/routes" // Import the routes package
    "github.com/rs/cors"  // Import the CORS package
)

func main() {
    // Get the MongoDB URI from the environment variable
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI environment variable not set")
    }

    // Connect to MongoDB database
    err := models.ConnectDB(mongoURI)
    if err != nil {
        log.Fatalf("Could not connect to MongoDB: %s\n", err)
    }

    // Initialize routes for the to-do service
    router := routes.SetupRoutes()

    // Set up CORS options for frontend origin
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:5173"}, // Replace with the production frontend origin
        AllowCredentials: true,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
    })

    // Use the CORS middleware for the router
    handler := c.Handler(router)

    // Start the server and listen on port 8082
    log.Println("To-Do service is running on :8082...")
    if err := http.ListenAndServe(":8082", handler); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}