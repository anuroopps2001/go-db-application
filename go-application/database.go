package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient interface {
	Ready() bool
	RunMigration() error
}

type Client struct {
	db *gorm.DB
}

func (c Client) Ready() bool {
	var ready string

	result := c.db.Raw("SELECT 1 as ready").Scan(&ready)
	if result.Error != nil {
		return false
	}

	return ready == "1"
}

func (c Client) RunMigration() error {
	if !c.Ready() {
		return fmt.Errorf("database is not ready")
	}

	return c.db.AutoMigrate(&User{})
}

func NewDBClient() (Client, error) {

	// 🔹 Read env vars
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbServer := os.Getenv("DB_SERVER") // 🔥 REQUIRED for Azure

	// 🔹 Validate required env vars
	if dbHost == "" || dbUsername == "" || dbPassword == "" || dbName == "" || dbPort == "" || dbServer == "" {
		return Client{}, fmt.Errorf("missing required database environment variables")
	}

	// 🔹 Convert port
	databasePort, err := strconv.Atoi(dbPort)
	if err != nil {
		return Client{}, fmt.Errorf("invalid DB port: %v", err)
	}

	// 🔥 Azure requires: username@server
	fullUsername := fmt.Sprintf("%s@%s", dbUsername, dbServer)

	// 🔹 Connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=require",
		dbHost,
		fullUsername,
		dbPassword,
		dbName,
		databasePort,
	)

	var db *gorm.DB

	// 🔥 Retry logic (important for Kubernetes startup)
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to database successfully")
			break
		}

		log.Printf("DB connection failed (%d/%d): %v", i+1, maxRetries, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		return Client{}, fmt.Errorf("failed to connect to DB after retries: %v", err)
	}

	return Client{db: db}, nil
}
