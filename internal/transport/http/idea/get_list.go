package idea

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/yazmeyaa/virtual_ideas_book/internal/common"
	"github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea/dto/response"
)

// Get implements IdeaHandler.
func (i *ideaHandler) Get(w http.ResponseWriter, r *http.Request) {
	offsetQ := r.URL.Query().Get("offset")
	limitQ := r.URL.Query().Get("limit")
	offset, _ := strconv.Atoi(offsetQ)
	limit, _ := strconv.Atoi(limitQ)

	list, err := i.ideaService.GetIdeasList(idea.GetIdeaListOptions{
		ListOptions: common.ListOptions{
			Limit:  limit,
			Offset: offset,
		},
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("Internal error: cannot get idea list from DB: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	res := response.GetIdeaListResponseDTO{}

	for _, _idea := range list {
		idea := response.GetIdeaByIdResponse{}
		idea.FromModel(_idea)
		res.Ideas = append(res.Ideas, idea)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
