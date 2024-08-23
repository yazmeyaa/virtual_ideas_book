package ideacomment

type CreateIdeaCommentDTO struct {
	IdeaID  uint
	UserID  uint
	Message string
}

type UpdateMessageDTO struct {
	Message string
}
