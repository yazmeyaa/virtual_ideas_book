package idea

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/middleware"
)

// Delete implements IdeaHandler.
func (i *ideaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, has := vars["id"]

	ideaId, err := strconv.Atoi(id)
	ideaIdAsUint := uint(ideaId)

	if err != nil || !has {
		http.Error(w, "Idea ID is not a number!", http.StatusBadRequest)
		return
	}

	userIdAny := r.Context().Value(middleware.USER_ID_KEY)
	userId, ok := userIdAny.(uint)

	if !ok {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	idea, err := i.ideaService.GetById(ideaIdAsUint)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	if idea.UserID != userId {
		http.Error(w, "You cannot delete this idea.", http.StatusMethodNotAllowed)
		return
	}

	i.ideaService.DeleteIdea(ideaIdAsUint)

	w.WriteHeader(http.StatusOK)
}
