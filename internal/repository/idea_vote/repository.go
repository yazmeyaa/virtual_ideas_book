package ideavote

import "github.com/yazmeyaa/virtual_ideas_book/internal/models"

type IdeaVoteRepository interface {
	// Create idea vote
	Create(value CreateIdeaVoteDTO) (models.IdeaVote, error)
	// Update idea vote
	Update(ideaVoteId uint, value UpdateIdeaVoteDTO) (models.IdeaVote, error)
	// Delete idea vote
	Delete(ideaVoteId uint) error
	// Get all idea votes.
	GetIdeaVotes(ideaId uint) ([]models.IdeaVote, error)
	// Get only positive or negative idea votes and count them.
	GetSpecificIdeaVotes(ideaId uint, positive bool) (votes []models.IdeaVote, err error)
}
