package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {

	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	db, err := NewDBClient()
	if err != nil {
		slog.Error("db init failed", "error", err.Error())
		return
	}

	err = db.RunMigration()
	if err != nil {
		slog.Error("migration failed", "error", err.Error())
		return
	}

	blobClient, err := NewBlobClient()
	if err != nil {
		slog.Error("blob init failed", "error", err.Error())
		return
	}

	// ✅ NEW: Kafka Producer
	producer := NewKafkaProducer()

	service := NewServer(db, blobClient, producer)

	slog.Info("starting server", "port", 8080)

	log.Fatal(service.Start())
}
