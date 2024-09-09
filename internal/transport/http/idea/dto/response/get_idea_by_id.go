package response

import (
	"time"

	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
)

type GetIdeaByIdResponse struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Content     string           `json:"content"`
	UserID      uint             `json:"userId"`
	Tags        []GetIdeaByIdTag `json:"tags"`
	CreatedAt   time.Time        `json:"createdAt"`
}

type GetIdeaByIdTag struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *GetIdeaByIdResponse) FromModel(model models.Idea) {
	r.ID = model.ID
	r.Title = model.Title
	r.Description = model.Description
	r.Content = model.Content
	r.UserID = model.UserID
	r.CreatedAt = model.Model.CreatedAt

	tags := make([]GetIdeaByIdTag, 0, len(model.Tags))

	for _, tag := range model.Tags {
		newTag := GetIdeaByIdTag{
			Id:          tag.ID,
			Name:        tag.Name,
			Description: tag.Description,
		}
		tags = append(tags, newTag)
	}

	r.Tags = tags
}
