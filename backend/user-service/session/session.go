package session

import (
    "net/http"
    "os"
    "github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
    // Retrieve the secret key from environment variable
    secretKey := os.Getenv("SESSION_SECRET")
    if secretKey == "" {
        panic("SESSION_SECRET environment variable not set")
    }
    store = sessions.NewCookieStore([]byte(secretKey))
}

func SetSession(w http.ResponseWriter, r *http.Request, email string) {
    session, _ := store.Get(r, "session")
    session.Values["email"] = email
    session.Save(r, w)
}

func GetSession(r *http.Request) string {
    session, _ := store.Get(r, "session")
    if email, ok := session.Values["email"].(string); ok {
        return email
    }
    return ""
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session")
    session.Options.MaxAge = -1
    session.Save(r, w)
}
