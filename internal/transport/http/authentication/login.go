package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/authentication/dto/request"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/authentication/dto/response"
)

// Login implements AuthenticationHandler.
func (a *authenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	var data request.LoginUserDTO

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&data); err != nil {
		a.log.Error("Error decoding request body", "error", err.Error())
		http.Error(w, "Invalid request payload", http.StatusUnprocessableEntity)
		return
	}

	if err := a.validate.Struct(data); err != nil {
		a.log.Error("Validation error", "error", err.Error())
		http.Error(w, "Validation failed", http.StatusUnprocessableEntity)
		return
	}

	var user models.User
	var err error

	if data.Username != "" {
		user, err = a.userService.GetUserByUsername(data.Username)
	} else if data.Email != "" {
		user, err = a.userService.GetUserByEmail(data.Email)
	} else {
		http.Error(w, "Username or Email must be provided", http.StatusBadRequest)
		return
	}

	if err != nil {
		a.log.Error("Error fetching user", "error", err.Error())
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	isPasswordValid := a.authService.ValidatePassword(user, data.Password)
	if !isPasswordValid {
		a.log.Error("Authentication error: not valid password")
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := a.authService.CreateToken(user)

	if err != nil {
		a.log.Error("Error while create token", "error", err.Error())
		http.Error(w, "Cannot create token", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.LoginUserResponse{
		UserID: user.ID,
	})
}
