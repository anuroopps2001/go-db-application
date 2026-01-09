package main

import "testing"

func TestDBConnection(t *testing.T) {
	client, err := NewDBClient()

	// t.Fatal("I was executed")
	if err != nil {
		t.Fatal("I was executed")
	}

	if !client.Ready() {
		t.Fatal("database is not reachable")
	}
}
