package idea

import "github.com/yazmeyaa/virtual_ideas_book/internal/models"

type CreateIdeaDTO struct {
	Title       string
	Description string
	Content     string
	UserID      uint
	Tags        []models.Tag
}

type UpdateIdeaDTO struct {
	Title       string
	Description string
	Content     string
	Tags        []models.Tag
}
