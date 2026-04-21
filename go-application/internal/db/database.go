package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go-application/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	db *gorm.DB
}

func (c Client) Ready(ctx context.Context) bool {
	sqlDB, err := c.db.DB()
	if err != nil {
		return false
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return sqlDB.PingContext(ctx) == nil
}

func (c Client) RunMigration() error {
	return c.db.AutoMigrate(&models.User{})
}

func (c Client) First(ctx context.Context, dest interface{}, id int) error {
	return c.db.WithContext(ctx).First(dest, id).Error
}

func (c Client) Save(ctx context.Context, value interface{}) error {
	return c.db.WithContext(ctx).Save(value).Error
}

func (c Client) Find(ctx context.Context, dest interface{}) error {
	return c.db.WithContext(ctx).Find(dest).Error
}

func (c Client) Delete(ctx context.Context, value interface{}, id int) error {
	return c.db.WithContext(ctx).Delete(value, id).Error
}

func (c Client) Create(ctx context.Context, value interface{}) error {
	return c.db.WithContext(ctx).Create(value).Error
}

func NewDBClient() (Client, error) {

	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbHost == "" || dbUsername == "" || dbPassword == "" || dbName == "" || dbPort == "" {
		return Client{}, fmt.Errorf("missing DB env vars")
	}

	port, err := strconv.Atoi(dbPort)
	if err != nil {
		return Client{}, err
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=require",
		dbHost, dbUsername, dbPassword, dbName, port,
	)

	var gormDB *gorm.DB

	for i := 0; i < 5; i++ {
		gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to DB")
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return Client{}, err
	}

	return Client{db: gormDB}, nil
}
