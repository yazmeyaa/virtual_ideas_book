package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/yazmeyaa/virtual_ideas_book/internal/repository/user"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/authentication/dto/request"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/authentication/dto/response"
)

func (a *authenticationHandler) Register(w http.ResponseWriter, r *http.Request) {
	userRequest := request.RegisterUserDTO{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		a.log.Error("Error decoding request body", "error", err.Error())
		http.Error(w, "Invalid request payload", http.StatusUnprocessableEntity)
		return
	}

	if err := a.validate.Struct(userRequest); err != nil {
		a.log.Error("Validation error", "error", err.Error())
		http.Error(w, "Validation failed", http.StatusUnprocessableEntity)
		return
	}

	user, err := a.userService.CreateUser(user.CreateUserDTO{
		Username: userRequest.Username,
		Password: userRequest.Password,
	})

	if err != nil {
		a.log.Error("Error while register user", "error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.RegisterResponse{
		UserID: user.ID,
	})
}
