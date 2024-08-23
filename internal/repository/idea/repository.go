package idea

import (
	"github.com/yazmeyaa/virtual_ideas_book/internal/common"
	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
)

type GetIdeaListOptions struct {
	common.ListOptions
	OrderBy string
}
type IdeaRepository interface {
	Create(idea CreateIdeaDTO) (models.Idea, error)
	Delete(ideaId uint) error
	Update(ideaId uint, idea UpdateIdeaDTO) (models.Idea, error)
	GetById(ideaId uint) (models.Idea, error)
	GetList(options common.ListOptions) ([]models.Idea, error)
	AddTagToIdea(idea *models.Idea, tag models.Tag) error
	RemoveTagFromIdea(idea *models.Idea, tag models.Tag) error
}
