package request

type UpdateIdeaRequestDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Tags        []uint `json:"tags"`
}
