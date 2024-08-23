package models

import "gorm.io/gorm"

type IdeaComment struct {
	gorm.Model
	IdeaID  uint
	UserID  uint
	Message string
}
