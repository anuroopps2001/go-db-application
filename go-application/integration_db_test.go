package main

// This is an integration test and should be ran with
// go test -tags=integration ./...
import (
	"os"
	"testing"
)

func TestDBConnection(t *testing.T) {

	if os.Getenv("DB_HOST") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_USERNAME") == "" ||
		os.Getenv("DB_PASSWORD") == "" ||
		os.Getenv("DB_NAME") == "" {
		t.Skip("Skipping DB integration test: DB env vars not set")
	}

	client, err := NewDBClient()

	// t.Fatal("I was executed")
	if err != nil {
		t.Fatal(err)
	}

	if !client.Ready() {
		t.Fatal("database is not reachable")
	}
}
