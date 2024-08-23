package user

import "github.com/yazmeyaa/virtual_ideas_book/internal/models"

type UserRepository interface {
	Create(user CreateUserDTO) (models.User, error)
	Delete(userId uint) error
	GetById(userId uint) (models.User, error)
	GetByUsername(username string) (models.User, error)
	GetByEmail(email string) (models.User, error)
	Update(userId uint, user UpdateUserDTO) (models.User, error)
}
