package idea

import (
	"github.com/yazmeyaa/virtual_ideas_book/internal/common"
	models "github.com/yazmeyaa/virtual_ideas_book/internal/models"
	idearepo "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea"
	ideacomment "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea_comment"
)

type IdeaService interface {
	GetIdeasList(optinos idearepo.GetIdeaListOptions) ([]models.Idea, error)
	GetById(id uint) (models.Idea, error)
	CreateIdea(dto idearepo.CreateIdeaDTO) (models.Idea, error)
	DeleteIdea(ideaId uint) error
	UpdateIdea(ideaId uint, dto idearepo.UpdateIdeaDTO) (models.Idea, error)
	AddTag(ideaId uint, tagId uint) error
	RemoveTag(ideaId uint, tagId uint) error
	AddComment(dto ideacomment.CreateIdeaCommentDTO) (models.IdeaComment, error)
	GetComments(ideaId uint, options common.ListOptions) ([]models.IdeaComment, error)
	GetTagById(tagId uint) (models.Tag, error)
}
