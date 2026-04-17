package events

import "time"

type UserCreatedEvent struct {
	Event  string    `json:"event"`
	UserID uint      `json:"user_id"`
	Email  string    `json:"email"`
	Name   string    `json:"name"`
	Time   time.Time `json:"time"`
}
