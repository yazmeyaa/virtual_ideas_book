package idea

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea/dto/response"
)

// GetById implements IdeaHandler.
func (i *ideaHandler) GetById(w http.ResponseWriter, r *http.Request) {
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

	idea, err := i.ideaService.GetById(uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := response.GetIdeaByIdResponse{}
	response.FromModel(idea)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
