package idea

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea/dto/request"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/middleware"
)

// Create implements IdeaHandler.
func (i *ideaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var data request.CreateIdeaRequestDTO

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&data); err != nil {
		i.log.Error("Error decoding request body", "error", err.Error())
		http.Error(w, "Invalid request payload", http.StatusUnprocessableEntity)
		return
	}

	if err := i.validate.Struct(data); err != nil {
		i.log.Error("Validation error", "error", err.Error())
		http.Error(w, "Validation failed", http.StatusUnprocessableEntity)
		return
	}

	userIdAny := r.Context().Value(middleware.USER_ID_KEY)
	userId, ok := userIdAny.(uint)

	if !ok {
		fmt.Printf("\n\nuserid: %+v\n\n", userIdAny)
		i.log.Error("Error while getting userId from middleware context", "Context value", userIdAny)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	tags := []models.Tag{}

	for _, tagId := range data.Tags {
		tag, err := i.ideaService.GetTagById(tagId)
		if err != nil {
			continue
		}
		tags = append(tags, tag)
	}

	idea, err := i.ideaService.CreateIdea(idea.CreateIdeaDTO{
		Title:       data.Title,
		Description: data.Description,
		Content:     data.Content,
		UserID:      uint(userId),
		Tags:        tags,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(idea)
}
