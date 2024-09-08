package response

import (
	"time"

	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
)

type CreateIdeaResponseDTO struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Content     string       `json:"content"`
	UserID      string       `json:"userId"`
	Tags        []models.Tag `json:"tags"`
	CraetedAt   time.Time    `json:"createdAt"`
}
