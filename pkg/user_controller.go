package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/IceWhaleTech/CasaOS/model"
)

var users []model.UserInfo

// CreateUser handles the creation of a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.UserInfo
    json.NewDecoder(r.Body).Decode(&user)

    if err := user.HashPassword(); err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    // Save user to the database (pseudo-code)
    users = append(users, user)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// GetUser handles fetching a user by username
func GetUser(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    for _, user := range users {
        if user.Username == username {
            json.NewEncoder(w).Encode(user)
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}

// DeleteUser handles deleting a user by username
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    for i, user := range users {
        if user.Username == username {
            users = append(users[:i], users[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }
    http.Error(w, "User not found", http.StatusNotFound)
}