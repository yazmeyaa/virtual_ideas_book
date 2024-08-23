package models

import "gorm.io/gorm"

type Idea struct {
	gorm.Model
	Title       string
	Description string
	Content     string
	UserID      uint
	Tags        []Tag `gorm:"many2many:idea_tags;"`
}
