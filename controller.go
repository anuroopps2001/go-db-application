package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *MuxServer) addUser(w http.ResponseWriter, r *http.Request) {
	var userData Userparam

	var user User

	json.NewDecoder(r.Body).Decode(&userData)

	user.Name = userData.Name

	user.Email = userData.Email

	user.Age = userData.Age

	s.db.Create(&user)

	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (s *MuxServer) listUsers(w http.ResponseWriter, r *http.Request) {
	var users User

	s.db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (s *MuxServer) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userData Userparam
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var user User
	if err := s.db.First(&user, userId).Error; err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	// Partial update
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

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User updated successfully",
	})
}

func (s *MuxServer) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var user User

	if err := s.db.First(&user, userId).Error; err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	if err := s.db.Delete(&user).Error; err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User deleted successfully",
	})
}
