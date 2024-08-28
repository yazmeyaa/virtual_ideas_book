package authentication

import "github.com/yazmeyaa/virtual_ideas_book/internal/models"

type AuthenticationService interface {
	CreateToken(user models.User) (string, error)
	ValidateToken(token string) (bool, error)
	ValidatePassword(user models.User, password string) bool
	HashPassword(password string) (string, error)
	RefreshToken(oldToken string) (string, error)
}
