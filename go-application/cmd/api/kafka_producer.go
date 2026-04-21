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
	topic  string
}

// 1. Initialize with the environment variable
func NewKafkaProducer() *KafkaProducer {

	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	if topic == "" {
		slog.Error("KAFKA_TOPIC not set")
		os.Exit(1)
	}

	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Balancer: &kafka.LeastBytes{},
		// ❌ DO NOT set Topic here
	}

	return &KafkaProducer{
		writer: writer,
		topic:  topic,
	}
}

// 2. Publish without worrying about the topic name
func (p *KafkaProducer) Publish(ctx context.Context, message interface{}) error {

	msgBytes, err := json.Marshal(message)
	if err != nil {
		slog.Error("failed to marshal kafka message", "error", err)
		return err
	}

	slog.Info("publishing kafka message",
		"topic", p.topic,
	)

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Topic: p.topic, // ✔ ONLY here
		Value: msgBytes,
	})

	if err != nil {
		slog.Error("kafka publish failed",
			"topic", p.topic,
			"error", err,
		)
		return err
	}

	slog.Info("kafka publish success",
		"topic", p.topic,
	)

	return nil
}
