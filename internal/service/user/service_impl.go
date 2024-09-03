package user_service

import (
	models "github.com/yazmeyaa/virtual_ideas_book/internal/models"
	user_repository "github.com/yazmeyaa/virtual_ideas_book/internal/repository/user"
	auth_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/authentication"
)

type userService struct {
	userRepo    user_repository.UserRepository
	authService auth_service.AuthenticationService
}

// CreateUser implements UserService.
func (u *userService) CreateUser(dto user_repository.CreateUserDTO) (models.User, error) {
	hash, err := u.authService.HashPassword(dto.Password)
	if err != nil {
		return models.User{}, err
	}
	dto.Password = hash
	return u.userRepo.Create(dto)
}

// DeleteUser implements UserService.
func (u *userService) DeleteUser(id uint) error {
	return u.userRepo.Delete(id)
}

// GetUserByEmail implements UserService.
func (u *userService) GetUserByEmail(email string) (models.User, error) {
	return u.userRepo.GetByEmail(email)
}

// GetUserById implements UserService.
func (u *userService) GetUserById(id uint) (models.User, error) {
	return u.userRepo.GetById(id)
}

// GetUserByUsername implements UserService.
func (u *userService) GetUserByUsername(username string) (models.User, error) {
	return u.userRepo.GetByUsername(username)
}

func New(
	userRepo user_repository.UserRepository,
	authService auth_service.AuthenticationService,
) UserService {
	return &userService{
		userRepo:    userRepo,
		authService: authService,
	}
}
