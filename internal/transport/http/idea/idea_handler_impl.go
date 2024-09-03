package idea

import (
	"net/http"

	"github.com/yazmeyaa/virtual_ideas_book/internal/service/idea"
)

type ideaHandler struct {
	ideaService *idea.IdeaService
}

// Create implements IdeaHandler.
func (i *ideaHandler) Create(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Delete implements IdeaHandler.
func (i *ideaHandler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Get implements IdeaHandler.
func (i *ideaHandler) Get(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetById implements IdeaHandler.
func (i *ideaHandler) GetById(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Update implements IdeaHandler.
func (i *ideaHandler) Update(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func New(ideaService idea.IdeaService) IdeaHandler {
	return &ideaHandler{}
}
