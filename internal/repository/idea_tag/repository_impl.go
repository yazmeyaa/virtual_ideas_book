package ideatag

import (
	"log/slog"

	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"gorm.io/gorm"
)

type ideaTagRepository struct {
	db  *gorm.DB
	log *slog.Logger
}

func (i *ideaTagRepository) GetById(id uint) (models.Tag, error) {
	tag := models.Tag{}

	if err := i.db.First(&tag, id).Error; err != nil {
		i.log.Error("[IdeaTagRepository.GetById] Error while retrieving tag by ID", "tagId", id, "error", err.Error())
		return models.Tag{}, err
	}

	i.log.Info("[IdeaTagRepository.GetById] Tag retrieved successfully", "tagId", id)
	return tag, nil
}

func (i *ideaTagRepository) GetByName(name string) (models.Tag, error) {
	tag := models.Tag{}

	if err := i.db.Where("name = ?", name).First(&tag).Error; err != nil {
		i.log.Error("[IdeaTagRepository.GetByName] Error while retrieving tag by name", "name", name, "error", err.Error())
		return models.Tag{}, err
	}

	i.log.Info("[IdeaTagRepository.GetByName] Tag retrieved successfully", "name", name)
	return tag, nil
}

func (i *ideaTagRepository) GetSpecificIdeaTags(ideaId uint) ([]models.Tag, error) {
	var idea models.Idea

	if err := i.db.Preload("Tags").First(&idea, ideaId).Error; err != nil {
		i.log.Error("[IdeaTagRepository.GetSpecificIdeaTags] Error while retrieving tags for idea", "ideaId", ideaId, "error", err.Error())
		return nil, err
	}

	i.log.Info("[IdeaTagRepository.GetSpecificIdeaTags] Tags retrieved successfully for idea", "ideaId", ideaId)
	return idea.Tags, nil
}

func (i *ideaTagRepository) Create(dto CreateTagDTO) (models.Tag, error) {
	tag := models.Tag{
		Name:        dto.Name,
		Description: dto.Description,
	}

	if err := i.db.Create(&tag).Error; err != nil {
		i.log.Error("[IdeaTagRepository.Create] Error while creating tag", "name", dto.Name, "error", err.Error())
		return models.Tag{}, err
	}

	i.log.Info("[IdeaTagRepository.Create] Tag created successfully", "tagId", tag.ID)
	return tag, nil
}

func (i *ideaTagRepository) Delete(id uint) error {
	var tag models.Tag
	if err := i.db.First(&tag, id).Error; err != nil {
		i.log.Error("[IdeaTagRepository.Delete] Error while retrieving tag by ID for deletion", "tagId", id, "error", err.Error())
		return err
	}

	if err := i.db.Delete(&models.Tag{}, id).Error; err != nil {
		i.log.Error("[IdeaTagRepository.Delete] Error while deleting tag", "tagId", id, "error", err.Error())
		return err
	}

	i.log.Info("[IdeaTagRepository.Delete] Tag deleted successfully", "tagId", id)
	return nil
}

func (i *ideaTagRepository) Update(tagId uint, dto UpdateTagDTO) (models.Tag, error) {
	var foundTag models.Tag

	if err := i.db.First(&foundTag, tagId).Error; err != nil {
		i.log.Error("[IdeaTagRepository.Update] Error while retrieving tag by ID for update", "tagId", tagId, "error", err.Error())
		return models.Tag{}, err
	}

	if dto.Description != "" {
		foundTag.Description = dto.Description
	}

	if dto.Name != "" {
		foundTag.Name = dto.Name
	}

	if err := i.db.Save(&foundTag).Error; err != nil {
		i.log.Error("[IdeaTagRepository.Update] Error while updating tag", "tagId", tagId, "error", err.Error())
		return models.Tag{}, err
	}

	i.log.Info("[IdeaTagRepository.Update] Tag updated successfully", "tagId", tagId)
	return foundTag, nil
}

func New(db *gorm.DB, log *slog.Logger) IdeaTagRepository {
	return &ideaTagRepository{
		db:  db,
		log: log,
	}
}
