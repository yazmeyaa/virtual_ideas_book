package idea

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	idea_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/idea"
	user_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/user"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea/dto/response"
)

type ideaHandler struct {
	ideaService idea_service.IdeaService
	userService user_service.UserService
	validate    *validator.Validate
	log         *slog.Logger
}

// Get implements IdeaHandler.
func (i *ideaHandler) Get(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

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

// Update implements IdeaHandler.
func (i *ideaHandler) Update(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func New(
	ideaService idea_service.IdeaService,
	userService user_service.UserService,
	validate *validator.Validate,
	log *slog.Logger,
) IdeaHandler {
	return &ideaHandler{
		ideaService: ideaService,
		userService: userService,
		validate:    validate,
		log:         log,
	}
}
