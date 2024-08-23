package database

import (
	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Tag{},
		&models.Idea{},
		&models.IdeaComment{},
		&models.IdeaVote{},
	); err != nil {
		return err
	}
	return nil
}
