package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

// 1. Initialize with the environment variable
func NewKafkaProducer() *KafkaProducer {
	return &KafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(os.Getenv("KAFKA_BROKER")),
			Topic:    os.Getenv("KAFKA_TOPIC"), // Defined once here
			Balancer: &kafka.LeastBytes{},
		},
	}
}

// 2. Publish without worrying about the topic name
func (p *KafkaProducer) Publish(ctx context.Context, topic string, message interface{}) error {

	// serialize message
	msgBytes, err := json.Marshal(message)
	if err != nil {
		slog.Error("failed to marshal kafka message", "error", err)
		return err
	}

	slog.Info("publishing kafka message",
		"topic", topic,
		"payload", string(msgBytes),
	)

	// send message
	err = p.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Value: msgBytes,
	})

	if err != nil {
		slog.Error("kafka publish failed",
			"topic", topic,
			"error", err,
		)
		return err
	}

	slog.Info("kafka publish success",
		"topic", topic,
	)

	return nil
}
