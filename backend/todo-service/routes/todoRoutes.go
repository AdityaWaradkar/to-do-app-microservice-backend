package routes

import (
	"todo-service/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/todos", handlers.AddTodoHandler).Methods("POST")

	router.HandleFunc("/api/todos/{id}", handlers.EditTodoHandler).Methods("PUT")

	router.HandleFunc("/api/todos/{id}", handlers.DeleteTodoHandler).Methods("DELETE")

	router.HandleFunc("/api/todos/fetch", handlers.FetchTodosHandler).Methods("GET")

	return router
}
