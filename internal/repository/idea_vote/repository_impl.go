package ideavote

import (
	"errors"
	"log/slog"

	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"gorm.io/gorm"
)

type ideaVoteRepository struct {
	db  *gorm.DB
	log *slog.Logger
}

func (i *ideaVoteRepository) Create(value CreateIdeaVoteDTO) (models.IdeaVote, error) {
	if value.IdeaID == 0 || value.UserID == 0 {
		i.log.Error("[IdeaVoteRepository.Create] Some of the required fields are empty")
		return models.IdeaVote{}, errors.New("some of the required fields are empty")
	}

	ideaVote := models.IdeaVote{
		IdeaID:         value.IdeaID,
		UserID:         value.UserID,
		IsPositiveVote: value.IsPositiveVote,
	}

	if err := i.db.Create(&ideaVote).Error; err != nil {
		i.log.Error("[IdeaVoteRepository.Create] Error while creating idea vote", "error", err.Error())
		return models.IdeaVote{}, err
	}

	i.log.Info("[IdeaVoteRepository.Create] Idea vote created successfully", "voteId", ideaVote.ID)
	return ideaVote, nil
}

func (i *ideaVoteRepository) Delete(ideaVoteId uint) error {
	if err := i.db.Delete(&models.IdeaVote{}, ideaVoteId).Error; err != nil {
		i.log.Error("[IdeaVoteRepository.Delete] Error while deleting idea vote", "voteId", ideaVoteId, "error", err.Error())
		return err
	}

	i.log.Info("[IdeaVoteRepository.Delete] Idea vote deleted successfully", "voteId", ideaVoteId)
	return nil
}

func (i *ideaVoteRepository) GetIdeaVotes(ideaId uint) ([]models.IdeaVote, error) {
	var ideaVotes []models.IdeaVote

	if err := i.db.Where("IdeaID = ?", ideaId).Find(&ideaVotes).Error; err != nil {
		i.log.Error("[IdeaVoteRepository.GetIdeaVotes] Error while retrieving idea votes", "ideaId", ideaId, "error", err.Error())
		return nil, err
	}

	i.log.Info("[IdeaVoteRepository.GetIdeaVotes] Idea votes retrieved successfully", "ideaId", ideaId)
	return ideaVotes, nil
}

func (i *ideaVoteRepository) GetSpecificIdeaVotes(ideaId uint, positive bool) (votes []models.IdeaVote, err error) {
	if err := i.db.Where("IdeaID = ? AND IsPositiveVote = ?", ideaId, positive).Find(&votes).Error; err != nil {
		i.log.Error("[IdeaVoteRepository.GetSpecificIdeaVotes] Error while retrieving specific idea votes", "ideaId", ideaId, "positive", positive, "error", err.Error())
		return nil, err
	}

	i.log.Info("[IdeaVoteRepository.GetSpecificIdeaVotes] Specific idea votes retrieved successfully", "ideaId", ideaId, "positive", positive)
	return votes, nil
}

func (i *ideaVoteRepository) Update(ideaVoteId uint, value UpdateIdeaVoteDTO) (models.IdeaVote, error) {
	vote := models.IdeaVote{}

	if err := i.db.Where("id = ?", ideaVoteId).First(&vote).Error; err != nil {
		i.log.Error("[IdeaVoteRepository.Update] Error while retrieving idea vote by ID", "voteId", ideaVoteId, "error", err.Error())
		return models.IdeaVote{}, err
	}

	vote.IsPositiveVote = value.IsPositiveVote

	if err := i.db.Save(&vote).Error; err != nil {
		i.log.Error("[IdeaVoteRepository.Update] Error while updating idea vote", "voteId", ideaVoteId, "error", err.Error())
		return models.IdeaVote{}, err
	}

	i.log.Info("[IdeaVoteRepository.Update] Idea vote updated successfully", "voteId", ideaVoteId)
	return vote, nil
}

func New(db *gorm.DB, log *slog.Logger) IdeaVoteRepository {
	return &ideaVoteRepository{
		db:  db,
		log: log,
	}
}
