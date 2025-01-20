package handlers

import (
	"encoding/json"
	"net/http"
	"todo-service/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FetchTodosHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		http.Error(w, `{"error": "Missing user ID"}`, http.StatusBadRequest)
		return
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, `{"error": "Invalid User ID"}`, http.StatusBadRequest)
		return
	}

	todos, err := models.FetchTodos(userObjectID)
	if err != nil {
		http.Error(w, `{"error": "Error fetching todos"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		http.Error(w, `{"error": "Failed to encode todos"}`, http.StatusInternalServerError)
		return
	}
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.TodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	userID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		http.Error(w, `{"error": "Invalid User ID"}`, http.StatusBadRequest)
		return
	}

	todo := models.Todo{
		ID:          primitive.NewObjectID(),
		Title:       input.Title,
		Description: input.Description,
		Completed:   input.Completed,
		UserID:      userID,
	}

	if err := todo.Save(); err != nil {
		http.Error(w, `{"error": "Failed to add todo"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status": "success", "message": "Todo added successfully"}`))
}

func EditTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	todoID := vars["id"]

	if todoID == "" {
		http.Error(w, `{"error": "Missing todo ID"}`, http.StatusBadRequest)
		return
	}

	if _, err := primitive.ObjectIDFromHex(todoID); err != nil {
		http.Error(w, `{"error": "Invalid todo ID format"}`, http.StatusBadRequest)
		return
	}

	var input models.TodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, `{"error": "Invalid input"}`, http.StatusBadRequest)
		return
	}

	if err := models.UpdateTodo(todoID, input); err != nil {
		http.Error(w, `{"error": "Failed to edit todo"}`, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"status": "success", "message": "Todo updated successfully"}`))
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	todoID := vars["id"]
	if todoID == "" {
		http.Error(w, `{"error": "Missing todo ID"}`, http.StatusBadRequest)
		return
	}

	if _, err := primitive.ObjectIDFromHex(todoID); err != nil {
		http.Error(w, `{"error": "Invalid todo ID format"}`, http.StatusBadRequest)
		return
	}

	if err := models.DeleteTodo(todoID); err != nil {
		http.Error(w, `{"error": "Failed to delete todo"}`, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"status": "success", "message": "Todo deleted successfully"}`))
}
