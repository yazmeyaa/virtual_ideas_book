package ideacomment

import (
	"github.com/yazmeyaa/virtual_ideas_book/internal/common"
	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
)

type IdeaCommentRepository interface {
	Create(dto CreateIdeaCommentDTO) (models.IdeaComment, error)
	Delete(commentId uint) error
	DeleteByIdea(ideaId uint) error
	Update(commentId uint, dto UpdateMessageDTO) (models.IdeaComment, error)
	GetById(commentId uint) (models.IdeaComment, error)
	GetSpecificIdeaComments(ideaId uint, options common.ListOptions) ([]models.IdeaComment, error)
}
