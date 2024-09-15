package ideatag_service

import (
	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	ideatag "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea_tag"
)

type TagService interface {
	GetTagByID(id uint) (models.Tag, error)
	GetTagByName(name string) (models.Tag, error)
	GetTagsForIdea(ideaId uint) ([]models.Tag, error)
	CreateTag(dto ideatag.CreateTagDTO) (models.Tag, error)
	UpdateTag(tagId uint, dto ideatag.UpdateTagDTO) (models.Tag, error)
	DeleteTag(tagId uint) error
}
