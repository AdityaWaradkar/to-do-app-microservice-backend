package routes

import (
	"user-service/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Public routes
    router.HandleFunc("/api/user/register", handlers.RegisterHandler).Methods("POST")
    router.HandleFunc("/api/user/login", handlers.LoginHandler).Methods("POST")

    // Protected route
    // Replace SessionMiddleware with JWTMiddleware
   

    return router
}
