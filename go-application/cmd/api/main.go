package main

import (
	"go-application/internal/blob"
	"go-application/internal/db"
	"log"
	"log/slog"
	"os"
)

func main() {

	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	db, err := db.NewDBClient()
	if err != nil {
		slog.Error("db init failed", "error", err)
		return
	}

	if err := db.RunMigration(); err != nil {
		slog.Error("migration failed", "error", err)
		return
	}

	blobClient, err := blob.NewBlobClient()
	if err != nil {
		slog.Error("blob init failed", "error", err)
		return
	}

	producer := NewKafkaProducer()

	server := NewServer(db, blobClient, producer)

	slog.Info("starting server", "port", 8080)

	log.Fatal(server.Start())
}
