package handlers

import (
    "encoding/json"
    "net/http"
    "user-service/models"
    "user-service/session"
)

type RegisterInput struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var input RegisterInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    user := models.User{
        Username: input.Username,
        Email:    input.Email,
        Password: input.Password,
    }

    if err := user.Save(); err != nil {
        http.Error(w, "Failed to register user", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Registration successful"))
}

type LoginInput struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var input LoginInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Find user by email in the database
    user, err := models.FindUserByEmail(input.Email)
    if err != nil || user.CheckPassword(input.Password) != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Set session for the logged-in user (optional, based on your session handling approach)
    session.SetSession(w, r, user.Email)

    // Respond with user data including the userID
    response := map[string]interface{}{
        "message": "OK",
        "userID": user.ID.Hex(), // Convert ObjectId to string
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}


func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    session.ClearSession(w, r)
    w.Write([]byte("Logged out"))
}
