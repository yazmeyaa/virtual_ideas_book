package http

import (
	"log/slog"
	"net/http"

	user_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/user"
)

type authenticationHandler struct {
	userService user_service.UserService
	log         *slog.Logger
}

// Login implements AuthenticationHandler.
func (a *authenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Logout implements AuthenticationHandler.
func (a *authenticationHandler) Logout(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Register implements AuthenticationHandler.
func (a *authenticationHandler) Register(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func New(
	userService user_service.UserService,
	log *slog.Logger,
) AuthenticationHandler {
	return &authenticationHandler{
		userService: userService,
		log:         log,
	}
}
