package routes

import (
    "github.com/gorilla/mux"
    "user-service/handlers" // Import user handlers
)

// SetupRoutes initializes the router and sets the routes
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // User routes
    router.HandleFunc("/api/users/register", handlers.RegisterUser).Methods("POST")
    router.HandleFunc("/api/users/login", handlers.LoginUser).Methods("POST")
    router.HandleFunc("/api/users/profile/{id}", handlers.UpdateUserProfile).Methods("PUT")
    router.HandleFunc("/api/users/reset-password", handlers.ResetPassword).Methods("POST")
    router.HandleFunc("/api/users/delete/{id}", handlers.DeleteUser).Methods("DELETE")

    return router
}
