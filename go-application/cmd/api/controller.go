package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go-application/internal/events"
	"go-application/internal/models"

	"github.com/gorilla/mux"
)

// CREATE USER
func (s *MuxServer) addUser(w http.ResponseWriter, r *http.Request) {

	var userData models.Userparam
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if userData.Name == "" || userData.Email == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	user.Name = userData.Name
	user.Email = userData.Email
	user.Age = userData.Age

	// ✅ use abstraction (NOT s.Client.db)
	if err := s.Client.Create(r.Context(), &user); err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, "email exists", http.StatusConflict)
			return
		}
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	event := events.UserCreatedEvent{
		Event:  "user_created",
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
		Time:   time.Now(),
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := s.producer.Publish(ctx, "user-events", event)
		if err != nil {
			slog.Error("kafka publish failed", "error", err)
		}
	}()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// LIST USERS
func (s *MuxServer) listUsers(w http.ResponseWriter, r *http.Request) {

	var users []models.User

	if err := s.Client.Find(r.Context(), &users); err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET USER
func (s *MuxServer) getUser(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var user models.User

	if err := s.Client.First(r.Context(), &user, userId); err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// UPDATE USER
func (s *MuxServer) updateUser(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := s.Client.First(r.Context(), &user, userId); err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	var input models.Userparam
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Age != 0 {
		user.Age = input.Age
	}

	if err := s.Client.Save(r.Context(), &user); err != nil {
		http.Error(w, "update failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// DELETE USER
func (s *MuxServer) deleteUser(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	if err := s.Client.Delete(r.Context(), &models.User{}, userId); err != nil {
		http.Error(w, "delete failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "user deleted",
	})
}

// HEALTH CHECK (liveness)
func (s *MuxServer) health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// READINESS CHECK
func (s *MuxServer) ready(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if !s.Client.Ready(ctx) {
		http.Error(w, "db not ready", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ready"))
}

// UPLOAD IMAGE
func (s *MuxServer) uploadProfileImage(w http.ResponseWriter, r *http.Request) {

	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "file read error", http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("user/%d/%d-%s",
		userId,
		time.Now().Unix(),
		handler.Filename,
	)

	url, err := s.blob.Upload(r.Context(), fileName, buffer)
	if err != nil {
		http.Error(w, "upload failed", http.StatusInternalServerError)
		return
	}

	event := events.UploadEvent{
		UserID:   userId,
		FileName: fileName,
		FileURL:  url,
		Time:     time.Now(),
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := s.producer.Publish(ctx, "upload-events", event)
		if err != nil {
			slog.Error("kafka publish failed", "error", err)
		}
	}()

	json.NewEncoder(w).Encode(map[string]string{
		"url": url,
	})
}
