package middleware

import (
    "net/http"
    "strings"
    "user-service/jwt" // Import the jwt package
)

// JWTMiddleware function to check if JWT is valid
func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get the JWT token from the Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
            return
        }

        // Split the token from the 'Bearer' prefix
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
            return
        }

        // Validate the token using the ValidateToken function from the jwt package
        token, err := jwt.ValidateToken(parts[1])
        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Pass the request to the next handler if token is valid
        next.ServeHTTP(w, r)
    })
}
