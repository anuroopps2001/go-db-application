package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer() *KafkaProducer {
	// Read the broker and default topic from environment variables
	broker := os.Getenv("KAFKA_BROKER")
	defaultTopic := os.Getenv("KAFKA_TOPIC") // e.g., "upload-events"

	writer := &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    defaultTopic, // Setting this allows p.writer.WriteMessages to know where to go
		Balancer: &kafka.LeastBytes{},
		// Required for better reliability in production
		Async: false,
	}

	return &KafkaProducer{
		writer: writer,
	}
}

// Publish now checks if a topic is provided; if not, it uses the default from the writer
func (p *KafkaProducer) Publish(ctx context.Context, topic string, message interface{}) error {
	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// If topic is empty string, the writer uses its internal defaultTopic
	err = p.writer.WriteMessages(ctx,
		kafka.Message{
			Topic: topic,
			Value: bytes,
		},
	)

	if err != nil {
		log.Println("kafka publish error:", err)
		return err
	}

	return nil
}
