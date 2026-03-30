package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

	user.Name = userData.Name
	user.Email = userData.Email
	user.Age = userData.Age

	if err := s.db.Create(&user).Error; err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// List Users
func (s *MuxServer) listUsers(w http.ResponseWriter, r *http.Request) {

	var users []User

	if err := s.db.Find(&users).Error; err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	if err := s.db.First(&user, userId).Error; err != nil {
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

	if err := s.db.Save(&user).Error; err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]string{
		"message": "User updated successfully",
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Delete User
func (s *MuxServer) deleteUser(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	if err := s.db.Delete(&User{}, userId).Error; err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Health check (liveness probe)
func (s *MuxServer) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte("ok")); err != nil {
		log.Println("Write failed:", err)
	}
}

// Readiness check (DB connectivity)
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

	if _, err := w.Write([]byte("ready")); err != nil {
		log.Println("Write failed:", err)
	}
}
