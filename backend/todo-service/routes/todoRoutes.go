package routes

import (
	"todo-service/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Route to add a new todo
	router.HandleFunc("/todos", handlers.AddTodoHandler).Methods("POST")

	// Route to edit an existing todo (with ID as URL parameter)
	router.HandleFunc("/todos/{id}", handlers.EditTodoHandler).Methods("PUT")

	// Route to delete a todo (with ID as URL parameter)
	router.HandleFunc("/todos/{id}", handlers.DeleteTodoHandler).Methods("DELETE")

	// Route to fetch todos for a specific user (using query parameter for userID)
	router.HandleFunc("/todos/fetch", handlers.FetchTodosHandler).Methods("GET")

	return router
}
