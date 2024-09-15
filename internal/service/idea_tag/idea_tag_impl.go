package ideatag_service

import (
	"errors"

	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	ideatag "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea_tag"
)

type tagService struct {
	repo ideatag.IdeaTagRepository
}

func NewTagService(repo ideatag.IdeaTagRepository) TagService {
	return &tagService{
		repo: repo,
	}
}

func (s *tagService) GetTagByID(id uint) (models.Tag, error) {
	tag, err := s.repo.GetById(id)
	if err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func (s *tagService) GetTagByName(name string) (models.Tag, error) {
	tag, err := s.repo.GetByName(name)
	if err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func (s *tagService) GetTagsForIdea(ideaId uint) ([]models.Tag, error) {
	tags, err := s.repo.GetSpecificIdeaTags(ideaId)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (s *tagService) CreateTag(dto ideatag.CreateTagDTO) (models.Tag, error) {
	if dto.Name == "" {
		return models.Tag{}, errors.New("tag name cannot be empty")
	}

	tag, err := s.repo.Create(dto)
	if err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func (s *tagService) UpdateTag(tagId uint, dto ideatag.UpdateTagDTO) (models.Tag, error) {
	tag, err := s.repo.Update(tagId, dto)
	if err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func (s *tagService) DeleteTag(tagId uint) error {
	if err := s.repo.Delete(tagId); err != nil {
		return err
	}
	return nil
}
