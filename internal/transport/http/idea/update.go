package idea

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea/dto/request"
)

// Update implements IdeaHandler.
func (i *ideaHandler) Update(w http.ResponseWriter, r *http.Request) {
	data := request.UpdateIdeaRequestDTO{}
	vars := mux.Vars(r)
	ideaId, has := vars["id"]
	if !has {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(ideaId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		i.log.Error("Error while decoding request body")
		http.Error(w, "Cannot decode request body", http.StatusBadRequest)
		return
	}

	tags := make([]models.Tag, 0, len(data.Tags))

	if data.Tags != nil {

		for _, tagId := range data.Tags {
			tag, err := i.ideaService.GetTagById(tagId)
			if err != nil {
				continue
			}

			tags = append(tags, tag)
		}
	}

	updateDTO := idea.UpdateIdeaDTO{
		Title:       data.Title,
		Description: data.Description,
		Content:     data.Content,
		Tags:        tags,
	}
	updated, err := i.ideaService.UpdateIdea(uint(id), updateDTO)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		i.log.Error(err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}
