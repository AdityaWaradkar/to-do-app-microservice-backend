package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux" // Import gorilla/mux
    "user-service/models"      // Import user models
)

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    registeredUser, err := models.RegisterUser(user)
    if err != nil {
        // Return a JSON error response
        w.Header().Set("Content-Type", "application/json") // Set content type
        w.WriteHeader(http.StatusConflict)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }

    w.Header().Set("Content-Type", "application/json") // Set content type
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(registeredUser)
}

// LoginUser handles user authentication
func LoginUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    authenticatedUser, err := models.AuthenticateUser(user.Username, user.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(authenticatedUser)
}

// UpdateUserProfile handles profile updates
func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"] // Get user ID from URL
    var updatedUser models.User
    if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := models.UpdateUserProfile(id, updatedUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

// ResetPassword handles password resets
func ResetPassword(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Username    string `json:"username"`
        NewPassword string `json:"new_password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := models.ResetPassword(request.Username, request.NewPassword)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Password reset successful"))
}

// DeleteUser handles user account deletion
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"] // Get user ID from URL
    err := models.DeleteUser(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent) // 204 No Content
}
