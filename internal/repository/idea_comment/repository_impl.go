package ideacomment

import (
	"errors"
	"log/slog"

	"github.com/yazmeyaa/virtual_ideas_book/internal/common"
	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"gorm.io/gorm"
)

type ideaCommentRepository struct {
	db  *gorm.DB
	log *slog.Logger
}

func (i *ideaCommentRepository) DeleteByIdea(ideaId uint) error {
	err := i.db.Where("IdeaID = ?", ideaId).Delete(&models.IdeaComment{}).Error
	if err != nil {
		i.log.Error("[IdeaCommentRepository.DeleteByIdea] Error while delete ideas comments", "error", err.Error())
	}

	return err
}

func (i *ideaCommentRepository) Create(dto CreateIdeaCommentDTO) (models.IdeaComment, error) {
	if dto.IdeaID == 0 || dto.UserID == 0 || dto.Message == "" {
		return models.IdeaComment{}, errors.New("some of the required fields are empty")
	}

	comment := models.IdeaComment{
		IdeaID:  dto.IdeaID,
		UserID:  dto.UserID,
		Message: dto.Message,
	}
	if err := i.db.Create(&comment).Error; err != nil {
		i.log.Error("[IdeaCommentRepository.Create] Error while create ideas comment", "error", err.Error())
		return models.IdeaComment{}, err
	}

	return comment, nil
}

func (i *ideaCommentRepository) Delete(commentId uint) error {
	comment := models.IdeaComment{}
	if err := i.db.First(&comment, commentId).Error; err != nil {
		i.log.Error("[IdeaCommentRepository.Delete] Error while delete ideas comment", "error", err.Error())
		return err
	}

	return i.db.Delete(&models.IdeaComment{}, commentId).Error
}

func (i *ideaCommentRepository) GetById(commentId uint) (models.IdeaComment, error) {
	comment := models.IdeaComment{}

	if err := i.db.Where("id = ?", commentId).First(&comment).Error; err != nil {
		i.log.Error("[IdeaCommentRepository.GetById] Error while recieve idea comment by ID", "error", err.Error())
		return models.IdeaComment{}, err
	}

	return comment, nil
}

func (i *ideaCommentRepository) GetSpecificIdeaComments(ideaId uint, options common.ListOptions) ([]models.IdeaComment, error) {
	var comments []models.IdeaComment

	if err := i.db.Where("IdeaID = ?", ideaId).Limit(options.Limit).Offset(options.Offset).Find(&comments).Error; err != nil {
		i.log.Error("[IdeaCommentRepository.GetSpecificIdeaComments] Error while recieve ideas comment by idea ID", "error", err.Error())
		return nil, err
	}

	return comments, nil
}

func (i *ideaCommentRepository) Update(commentId uint, dto UpdateMessageDTO) (models.IdeaComment, error) {
	comment := models.IdeaComment{}
	if err := i.db.Where("id = ?", commentId).First(&comment).Error; err != nil {
		i.log.Error("[IdeaCommentRepository.Update] Error while get idea before update it", "error", err.Error())
		return models.IdeaComment{}, err
	}

	if dto.Message != "" {
		comment.Message = dto.Message
	}

	if err := i.db.Save(&comment).Error; err != nil {
		i.log.Error("[IdeaCommentRepository.Update] Error while update ideas comment", "error", err.Error())
		return models.IdeaComment{}, err
	}

	return comment, nil
}

func New(db *gorm.DB, log *slog.Logger) IdeaCommentRepository {
	return &ideaCommentRepository{
		db: db,
	}
}
