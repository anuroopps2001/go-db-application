package events

import "time"

type UserCreatedEvent struct {
	Event  string    `json:"event"`
	UserID uint      `json:"user_id"`
	Email  string    `json:"email"`
	Name   string    `json:"name"`
	Time   time.Time `json:"time"`
}

type UploadEvent struct {
	UserID   int       `json:"user_id"`
	FileName string    `json:"file_name"`
	FileURL  string    `json:"file_url"` // ✅ FIXED
	Time     time.Time `json:"time"`
}
