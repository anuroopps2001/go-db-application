package main

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"os" // Required to read environment variables

	"go-application/internal/events"

	"github.com/segmentio/kafka-go"
)

func main() {
	// 1. Get the broker address from the environment variable 'KAFKA_BROKER'
	// In your docker-compose, this is set to "kafka:9092"
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		// Fallback for local development outside of Docker
		broker = "localhost:9092"
	}

	slog.Info("Starting consumer", "broker", broker)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker}, // Using the dynamic address
		Topic:   "user-events",
		GroupID: "user-group",
	})

	// It's good practice to close the reader when the function exits
	defer reader.Close()

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
