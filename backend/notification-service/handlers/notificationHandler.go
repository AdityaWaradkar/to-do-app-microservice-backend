package handlers

import (
    "encoding/json"
    "net/http"
    "notification-service/models"
)

// SendNotification handles creating a new notification
func SendNotification(w http.ResponseWriter, r *http.Request) {
    var notification models.Notification
    if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    createdNotification, err := models.AddNotification(notification.Message)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdNotification)
}
