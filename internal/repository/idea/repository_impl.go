package idea

import (
	"errors"
	"log/slog"

	"github.com/yazmeyaa/virtual_ideas_book/internal/common"
	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"gorm.io/gorm"
)

type ideaRepository struct {
	db  *gorm.DB
	log *slog.Logger
}

func (r *ideaRepository) RemoveTagFromIdea(idea *models.Idea, tag models.Tag) error {
	tags := []models.Tag{}

	for _, t := range idea.Tags {
		if t.ID != tag.ID {
			tags = append(tags, t)
		}
	}

	idea.Tags = tags

	if err := r.db.Save(idea).Error; err != nil {
		r.log.Error("[IdeaRepository.RemoveTagFromIdea] Error while removing tag from idea", "ideaId", idea.ID, "tagId", tag.ID, "error", err.Error())
		return err
	}

	r.log.Info("[IdeaRepository.RemoveTagFromIdea] Tag removed from idea successfully", "ideaId", idea.ID, "tagId", tag.ID)
	return nil
}

func (r *ideaRepository) AddTagToIdea(idea *models.Idea, tag models.Tag) error {
	idea.Tags = append(idea.Tags, tag)

	if err := r.db.Save(idea).Error; err != nil {
		r.log.Error("[IdeaRepository.AddTagToIdea] Error while adding tag to idea", "ideaId", idea.ID, "tagId", tag.ID, "error", err.Error())
		return err
	}

	r.log.Info("[IdeaRepository.AddTagToIdea] Tag added to idea successfully", "ideaId", idea.ID, "tagId", tag.ID)
	return nil
}

func (r *ideaRepository) Create(idea CreateIdeaDTO) (models.Idea, error) {
	if idea.Content == "" || idea.Description == "" || idea.Title == "" {
		r.log.Error("[IdeaRepository.Create] Required fields are empty", "title", idea.Title, "userId", idea.UserID)
		return models.Idea{}, errors.New("some of the required fields are empty")
	}

	newIdea := models.Idea{
		Title:       idea.Title,
		Description: idea.Description,
		Content:     idea.Content,
		UserID:      idea.UserID,
		Tags:        idea.Tags,
	}

	if err := r.db.Create(&newIdea).Error; err != nil {
		r.log.Error("[IdeaRepository.Create] Error while creating idea", "title", idea.Title, "userId", idea.UserID, "error", err.Error())
		return models.Idea{}, err
	}

	r.log.Info("[IdeaRepository.Create] Idea created successfully", "ideaId", newIdea.ID)
	return newIdea, nil
}

func (r *ideaRepository) Delete(ideaId uint) error {
	if err := r.db.Where("id = ?", ideaId).Delete(&models.Idea{}).Error; err != nil {
		r.log.Error("[IdeaRepository.Delete] Error while deleting idea", "ideaId", ideaId, "error", err.Error())
		return err
	}

	r.log.Info("[IdeaRepository.Delete] Idea deleted successfully", "ideaId", ideaId)
	return nil
}

func (r *ideaRepository) GetById(ideaId uint) (models.Idea, error) {
	idea := models.Idea{}
	if err := r.db.Preload("Tags").Where("id = ?", ideaId).First(&idea).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Warn("[IdeaRepository.GetById] Idea not found", "ideaId", ideaId)
			return models.Idea{}, errors.New("idea not found")
		}
		r.log.Error("[IdeaRepository.GetById] Error while retrieving idea by ID", "ideaId", ideaId, "error", err.Error())
		return models.Idea{}, err
	}

	r.log.Info("[IdeaRepository.GetById] Idea retrieved successfully", "ideaId", ideaId)
	return idea, nil
}

func (r *ideaRepository) GetList(options common.ListOptions) ([]models.Idea, error) {
	var list []models.Idea

	limit := 20
	if options.Limit != 0 {
		limit = options.Limit
	}

	if err := r.db.Preload("Tags").Offset(options.Offset).Limit(limit).Find(&list).Error; err != nil {
		r.log.Error("[IdeaRepository.GetList] Error while retrieving ideas list", "error", err.Error())
		return nil, err
	}

	r.log.Info("[IdeaRepository.GetList] Ideas list retrieved successfully", "count", len(list))
	return list, nil
}

func (r *ideaRepository) Update(ideaId uint, idea UpdateIdeaDTO) (models.Idea, error) {
	foundIdea := models.Idea{}
	if err := r.db.Preload("Tags").Where("id = ?", ideaId).First(&foundIdea).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Warn("[IdeaRepository.Update] Idea not found", "ideaId", ideaId)
			return models.Idea{}, errors.New("idea not found")
		}
		r.log.Error("[IdeaRepository.Update] Error while retrieving idea by ID for update", "ideaId", ideaId, "error", err.Error())
		return models.Idea{}, err
	}

	if idea.Title != "" {
		foundIdea.Title = idea.Title
	}
	if idea.Description != "" {
		foundIdea.Description = idea.Description
	}
	if idea.Content != "" {
		foundIdea.Content = idea.Content
	}
	if len(idea.Tags) != len(foundIdea.Tags) {
		if err := r.db.Model(&foundIdea).Association("Tags").Replace(idea.Tags); err != nil {
			r.log.Error("[IdeaRepository.Update] Error while updating idea tags", "ideaId", ideaId, "error", err.Error())
			return models.Idea{}, err
		}
	}

	if err := r.db.Save(&foundIdea).Error; err != nil {
		r.log.Error("[IdeaRepository.Update] Error while updating idea", "ideaId", ideaId, "error", err.Error())
		return models.Idea{}, err
	}

	r.log.Info("[IdeaRepository.Update] Idea updated successfully", "ideaId", ideaId)
	return foundIdea, nil
}

func New(db *gorm.DB, logger *slog.Logger) IdeaRepository {
	return &ideaRepository{
		db:  db,
		log: logger,
	}
}
