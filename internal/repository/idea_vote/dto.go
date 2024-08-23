package ideavote

type CreateIdeaVoteDTO struct {
	UserID         uint
	IdeaID         uint
	IsPositiveVote bool
}

type UpdateIdeaVoteDTO struct {
	IsPositiveVote bool
}
