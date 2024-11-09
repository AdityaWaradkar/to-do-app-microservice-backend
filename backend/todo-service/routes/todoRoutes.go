package routes

import (
    "github.com/gorilla/mux"
    "todo-service/handlers"
)

// SetupRoutes initializes the router and sets the routes
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/api/todos", handlers.AddToDo).Methods("POST")
    router.HandleFunc("/api/todos/{id}", handlers.GetToDo).Methods("GET")
    router.HandleFunc("/api/todos/{id}", handlers.UpdateToDo).Methods("PUT")
    router.HandleFunc("/api/todos/{id}", handlers.DeleteToDo).Methods("DELETE")

    return router
}