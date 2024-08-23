package user_service

import (
	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"github.com/yazmeyaa/virtual_ideas_book/internal/repository/user"
)

type UserService interface {
	GetUserById(id uint) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	CreateUser(dto user.CreateUserDTO) (models.User, error)
	DeleteUser(id uint) error
}
