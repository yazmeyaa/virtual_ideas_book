package request

type CreateIdeaRequestDTO struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description" validate:"required"`
	Content     string `json:"content" validate:"required"`
	Tags        []uint `json:"tags"`
}
