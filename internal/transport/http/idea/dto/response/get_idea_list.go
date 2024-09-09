package response

type GetIdeaListResponseDTO struct {
	Ideas []GetIdeaByIdResponse `json:"ideas"`
}
