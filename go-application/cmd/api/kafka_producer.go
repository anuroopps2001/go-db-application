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

	writer := &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_BROKER")),
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{
		writer: writer,
	}
}

func (p *KafkaProducer) Publish(ctx context.Context, topic string, message interface{}) error {

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

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
