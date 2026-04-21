package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"time"

	"go-application/internal/db"
	"go-application/internal/events"
	"go-application/internal/models"

	"github.com/segmentio/kafka-go"
)

func main() {

	broker := os.Getenv("KAFKA_BROKER")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		GroupID: "worker-group",
		Topic:   "upload-events",
	})

	dbClient, _ := db.NewDBClient()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			continue
		}

		var event events.UploadEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			continue
		}

		// ✅ update DB only
		var user models.User
		if err := dbClient.First(&user, event.UserID); err != nil {
			continue
		}

		user.ProfileImage = event.FileURL

		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		dbClient.Save(&user)
		cancel()

		slog.Info("profile updated",
			"user_id", event.UserID,
			"url", event.FileURL,
		)
	}
}
