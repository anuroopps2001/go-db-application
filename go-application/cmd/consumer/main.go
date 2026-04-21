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
	if broker == "" {
		slog.Error("KAFKA_BROKER not set")
		return
	}

	slog.Info("starting consumer", "broker", broker)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		GroupID:  "worker-group",
		Topic:    "upload-events",
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	defer reader.Close()

	dbClient, err := db.NewDBClient()
	if err != nil {
		slog.Error("db init failed", "error", err)
		return
	}

	for {
		// ✅ use cancellable context for kafka read
		ctx := context.Background()

		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			slog.Error("failed to fetch message", "error", err)
			continue
		}

		var event events.UploadEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			slog.Error("failed to unmarshal event", "error", err)
			// commit bad message so it doesn't block
			_ = reader.CommitMessages(ctx, msg)
			continue
		}

		// ✅ DB operations with timeout context
		dbCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		var user models.User
		err = dbClient.First(dbCtx, &user, event.UserID)
		if err != nil {
			slog.Error("user fetch failed",
				"user_id", event.UserID,
				"error", err,
			)
			cancel()
			continue // do NOT commit → retry later
		}

		user.ProfileImage = event.FileURL

		err = dbClient.Save(dbCtx, &user)
		cancel()

		if err != nil {
			slog.Error("db update failed",
				"user_id", event.UserID,
				"error", err,
			)
			continue // do NOT commit → retry
		}

		// ✅ commit only after success
		if err := reader.CommitMessages(ctx, msg); err != nil {
			slog.Error("failed to commit message", "error", err)
			continue
		}

		slog.Info("profile updated",
			"user_id", event.UserID,
			"url", event.FileURL,
		)
	}
}
