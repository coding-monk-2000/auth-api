package storage

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func buildPostgresConnStr() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSL_MODE"),
	)
}

func InitDatabase() (AuthStore, error) {
	var db *gorm.DB
	var err error
	switch os.Getenv("DB_DRIVER") {
	case "postgres":
		db, err = gorm.Open(postgres.Open(buildPostgresConnStr()), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", os.Getenv("DB_DRIVER"))
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	store := NewGormStore(db)

	return store, nil
}
