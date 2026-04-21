package main

import (
	"context"
	"encoding/json"
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
func (p *KafkaProducer) Publish(ctx context.Context, message interface{}) error {
	bytes, _ := json.Marshal(message)
	return p.writer.WriteMessages(ctx, kafka.Message{
		Value: bytes, // The writer already knows where to go
	})
}
