package idea

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	idea_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/idea"
	user_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/user"
)

type ideaHandler struct {
	ideaService idea_service.IdeaService
	userService user_service.UserService
	validate    *validator.Validate
	log         *slog.Logger
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
