package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {

	// 🔥 JSON logging (required for Loki)
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

	service := NewServer(db)

	slog.Info("starting server", "port", 8080)

	log.Fatal(service.Start())
}
