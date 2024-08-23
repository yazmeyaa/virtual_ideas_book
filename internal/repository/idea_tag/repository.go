package ideatag

import "github.com/yazmeyaa/virtual_ideas_book/internal/models"

type IdeaTagRepository interface {
	Create(dto CreateTagDTO) (models.Tag, error)
	GetByName(name string) (models.Tag, error)
	GetById(id uint) (models.Tag, error)
	Delete(id uint) error
	Update(tagId uint, dto UpdateTagDTO) (models.Tag, error)
	GetSpecificIdeaTags(ideaId uint) ([]models.Tag, error)
}
