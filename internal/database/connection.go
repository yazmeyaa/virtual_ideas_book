package database

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(dsn string, log *slog.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to database: %s\n", "error", err.Error())
		return nil, err
	}

	return db, nil
}
