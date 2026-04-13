package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Create User
func (s *MuxServer) addUser(w http.ResponseWriter, r *http.Request) {

	var userData Userparam
	var user User

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if userData.Name == "" || userData.Email == "" {

		slog.Warn("invalid user input",
			"request_id", getRequestID(r.Context()),
			"name", userData.Name,
			"email", userData.Email,
		)

		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	user.Name = userData.Name
	user.Email = userData.Email
	user.Age = userData.Age

	err := observeDBWithContext(r.Context(), "create_user", func() error {
		return s.db.Create(&user).Error
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			http.Error(w, "email already exists", http.StatusConflict)
			return
		}
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

// List Users
func (s *MuxServer) listUsers(w http.ResponseWriter, r *http.Request) {

	var users []User

	err := observeDBWithContext(r.Context(), "list_users", func() error {
		return s.db.Find(&users).Error
	})

	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get User
func (s *MuxServer) getUser(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var user User

	err = observeDBWithContext(r.Context(), "get_user", func() error {
		return s.db.First(&user, userId).Error
	})

	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Update User
func (s *MuxServer) updateUser(w http.ResponseWriter, r *http.Request) {

	var userData Userparam

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var user User

	err = observeDBWithContext(r.Context(), "get_user", func() error {
		return s.db.First(&user, userId).Error
	})

	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	if userData.Name != "" {
		user.Name = userData.Name
	}
	if userData.Email != "" {
		user.Email = userData.Email
	}
	if userData.Age != 0 {
		user.Age = userData.Age
	}

	err = observeDBWithContext(r.Context(), "update_user", func() error {
		return s.db.Save(&user).Error
	})

	if err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User updated successfully",
	})
}

// Delete User
func (s *MuxServer) deleteUser(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	err = observeDBWithContext(r.Context(), "delete_user", func() error {
		return s.db.Delete(&User{}, userId).Error
	})

	if err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}

// Health check
func (s *MuxServer) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// Readiness check
func (s *MuxServer) ready(w http.ResponseWriter, _ *http.Request) {

	sqlDB, err := s.db.DB()
	if err != nil {
		http.Error(w, "db handle error", http.StatusServiceUnavailable)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		http.Error(w, "db not ready", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ready"))
}
