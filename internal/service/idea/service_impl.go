package idea

import (
	"log/slog"

	"github.com/yazmeyaa/virtual_ideas_book/internal/common"
	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	idea_repository "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea"
	ideacomment_repository "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea_comment"
	ideatag_repository "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea_tag"
	ideavote_repository "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea_vote"
)

type ideaService struct {
	commentRepo ideacomment_repository.IdeaCommentRepository
	ideaRepo    idea_repository.IdeaRepository
	tagRepo     ideatag_repository.IdeaTagRepository
	voteRepo    ideavote_repository.IdeaVoteRepository
	log         *slog.Logger
}

// GetIdeasList implements IdeaService.
func (i *ideaService) GetIdeasList(options idea_repository.GetIdeaListOptions) ([]models.Idea, error) {
	return i.ideaRepo.GetList(options.ListOptions)
}

// AddComment implements IdeaService.
func (i *ideaService) AddComment(dto ideacomment_repository.CreateIdeaCommentDTO) (models.IdeaComment, error) {
	return i.commentRepo.Create(dto)
}

// GetComments implements IdeaService.
func (i *ideaService) GetComments(ideaId uint, options common.ListOptions) ([]models.IdeaComment, error) {
	return i.commentRepo.GetSpecificIdeaComments(ideaId, options)
}

func (i *ideaService) AddTag(ideaId uint, tagId uint) error {
	tag, err := i.tagRepo.GetById(tagId)
	if err != nil {
		i.log.Error("[IdeaService.AddTag] Error while getting tag by ID", "tagId", tagId, "error", err.Error())
		return err
	}

	idea, err := i.ideaRepo.GetById(ideaId)
	if err != nil {
		i.log.Error("[IdeaService.AddTag] Error while getting idea by ID", "ideaId", ideaId, "error", err.Error())
		return err
	}

	for _, existingTag := range idea.Tags {
		if existingTag.ID == tag.ID {
			i.log.Info("[IdeaService.AddTag] Tag already exists for idea", "ideaId", ideaId, "tagId", tagId)
			return nil
		}
	}

	err = i.ideaRepo.AddTagToIdea(&idea, tag)
	if err != nil {
		i.log.Error("[IdeaService.AddTag] Error while adding tag to idea", "ideaId", ideaId, "tagId", tagId, "error", err.Error())
		return err
	}

	i.log.Info("[IdeaService.AddTag] Tag added to idea successfully", "ideaId", ideaId, "tagId", tagId)
	return nil
}

func (i *ideaService) RemoveTag(ideaId uint, tagId uint) error {
	tag, err := i.tagRepo.GetById(tagId)
	if err != nil {
		i.log.Error("[IdeaService.RemoveTag] Error while getting tag by ID", "tagId", tagId, "error", err.Error())
		return err
	}

	idea, err := i.ideaRepo.GetById(ideaId)
	if err != nil {
		i.log.Error("[IdeaService.RemoveTag] Error while getting idea by ID", "ideaId", ideaId, "error", err.Error())
		return err
	}

	err = i.ideaRepo.RemoveTagFromIdea(&idea, tag)
	if err != nil {
		i.log.Error("[IdeaService.RemoveTag] Error while removing tag from idea", "ideaId", ideaId, "tagId", tagId, "error", err.Error())
		return err
	}

	i.log.Info("[IdeaService.RemoveTag] Tag removed from idea successfully", "ideaId", ideaId, "tagId", tagId)
	return nil
}

func (i *ideaService) CreateIdea(dto idea_repository.CreateIdeaDTO) (models.Idea, error) {
	idea, err := i.ideaRepo.Create(dto)
	if err != nil {
		i.log.Error("[IdeaService.CreateIdea] Error while creating idea", "dto", dto, "error", err.Error())
		return models.Idea{}, err
	}

	i.log.Info("[IdeaService.CreateIdea] Idea created successfully", "ideaId", idea.ID)
	return idea, nil
}

func (i *ideaService) DeleteIdea(ideaId uint) error {
	err := i.ideaRepo.Delete(ideaId)
	if err != nil {
		i.log.Error("[IdeaService.DeleteIdea] Error while deleting idea", "ideaId", ideaId, "error", err.Error())
		return err
	}

	i.log.Info("[IdeaService.DeleteIdea] Idea deleted successfully", "ideaId", ideaId)
	return nil
}

func (i *ideaService) UpdateIdea(ideaId uint, dto idea_repository.UpdateIdeaDTO) (models.Idea, error) {
	idea, err := i.ideaRepo.Update(ideaId, dto)
	if err != nil {
		i.log.Error("[IdeaService.UpdateIdea] Error while updating idea", "ideaId", ideaId, "dto", dto, "error", err.Error())
		return models.Idea{}, err
	}

	i.log.Info("[IdeaService.UpdateIdea] Idea updated successfully", "ideaId", idea.ID)
	return idea, nil
}

func New(
	ideaRepo idea_repository.IdeaRepository,
	commentRepo ideacomment_repository.IdeaCommentRepository,
	tagRepo ideatag_repository.IdeaTagRepository,
	voteRepo ideavote_repository.IdeaVoteRepository,
	log *slog.Logger,
) IdeaService {
	return &ideaService{
		commentRepo: commentRepo,
		ideaRepo:    ideaRepo,
		tagRepo:     tagRepo,
		voteRepo:    voteRepo,
		log:         log,
	}
}
