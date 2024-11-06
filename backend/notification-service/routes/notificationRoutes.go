package routes

import (
    "github.com/gorilla/mux"
    "notification-service/handlers"
)

// SetupRoutes initializes the routes for the notification service.
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Route for sending notifications
    router.HandleFunc("/notifications", handlers.SendNotification).Methods("POST")

    return router
}
