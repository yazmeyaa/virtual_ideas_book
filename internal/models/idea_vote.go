package models

import "gorm.io/gorm"

type IdeaVote struct {
	gorm.Model
	UserID         uint
	IdeaID         uint
	IsPositiveVote bool
}
