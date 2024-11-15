package routes

import (
    "github.com/gorilla/mux"
    "net/http"
    "user-service/handlers"
    "user-service/middleware"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Public routes
    router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
    router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

    // Protected route
    router.Handle("/logout", middleware.SessionMiddleware(http.HandlerFunc(handlers.LogoutHandler))).Methods("POST")

    return router
}