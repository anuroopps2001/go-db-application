package main

// run with: go test -tags=integration ./...

import (
	"context"
	"os"
	"testing"
	"time"

	"go-application/internal/db"
)

func TestDBConnection(t *testing.T) {

	if os.Getenv("DB_HOST") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_USERNAME") == "" ||
		os.Getenv("DB_PASSWORD") == "" ||
		os.Getenv("DB_NAME") == "" {
		t.Skip("Skipping DB integration test: DB env vars not set")
	}

	client, err := db.NewDBClient()
	if err != nil {
		t.Fatal(err)
	}

	// ✅ use context (matches your DB layer now)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if !client.Ready(ctx) {
		t.Fatal("database is not reachable")
	}
}
