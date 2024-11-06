package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "todo-service/models"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// AddToDo handles creating a new to-do item
func AddToDo(w http.ResponseWriter, r *http.Request) {
    var todo models.ToDo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    createdToDo, err := models.AddToDo(todo)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdToDo)
}

// GetToDo handles retrieving a specific to-do item by ID
func GetToDo(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    todo, err := models.GetToDoByID(objectId)
    if err != nil {
        http.Error(w, "To-Do item not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(todo)
}

// UpdateToDo handles updating a to-do item by ID
func UpdateToDo(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    var updatedToDo models.ToDo
    if err := json.NewDecoder(r.Body).Decode(&updatedToDo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    todo, err := models.UpdateToDoByID(objectId, updatedToDo)
    if err != nil {
        http.Error(w, "To-Do item not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(todo)
}

// DeleteToDo handles deleting a to-do item by ID
func DeleteToDo(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    err = models.DeleteToDoByID(objectId)
    if err != nil {
        http.Error(w, "To-Do item not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
