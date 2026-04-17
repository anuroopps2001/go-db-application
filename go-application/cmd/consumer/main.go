package main

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"

	"go-application/internal/events"

	"github.com/segmentio/kafka-go"
)

func main() {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "user-events",
		GroupID: "user-group",
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("error reading message:", err)
			continue
		}

		var event events.UserCreatedEvent

		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			slog.Error("failed to parse event", "error", err)
			continue
		}

		slog.Info("user event received",
			"user_id", event.UserID,
			"email", event.Email,
			"name", event.Name,
		)
	}
}
