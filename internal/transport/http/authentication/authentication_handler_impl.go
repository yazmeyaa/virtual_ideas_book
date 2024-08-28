// Authentication HTTP-handler package
package authentication

import (
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	authentication_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/authentication"
	user_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/user"
)

type authenticationHandler struct {
	userService user_service.UserService
	authService authentication_service.AuthenticationService
	log         *slog.Logger
	validate    *validator.Validate
}

// Logout implements AuthenticationHandler.
func (a *authenticationHandler) Logout(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func New(
	userService user_service.UserService,
	authService authentication_service.AuthenticationService,
	log *slog.Logger,
	validate *validator.Validate,
) AuthenticationHandler {
	return &authenticationHandler{
		userService: userService,
		log:         log,
		validate:    validate,
		authService: authService,
	}
}
