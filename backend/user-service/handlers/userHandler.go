package handlers

import (
    "encoding/json"
    "net/http"
    "user-service/models"
    "user-service/jwt"  // Assuming you have a jwt package for token creation
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

    // Send a JSON response on successful registration
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Registration successful"})
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

    // Generate JWT token for the user
    token, err := jwt.GenerateToken(user.ID.Hex())  // Assuming GenerateToken is a function in your jwt package
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    // Respond with JWT token and userID
    response := map[string]interface{}{
        "message": "OK",
        "token":   token,  // Send the JWT token
        "userID":  user.ID.Hex(), // Send the userID in the response
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

