package user

import (
	"errors"
	"log/slog"

	"github.com/yazmeyaa/virtual_ideas_book/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db  *gorm.DB
	log *slog.Logger
}

func (u *userRepository) Create(userDTO CreateUserDTO) (models.User, error) {
	if userDTO.Password == "" || userDTO.Username == "" {
		u.log.Error("[UserRepository.Create] Some of required fields are empty")
		return models.User{}, errors.New("some of required fields are empty")
	}

	user := models.User{
		Username: userDTO.Username,
		Password: userDTO.Password,
		Email:    userDTO.Email,
	}

	var email string

	if user.Email != nil {
		email = *user.Email
	}

	existingUser := models.User{}
	if err := u.db.Where("email = ? OR username = ?", email, user.Username).First(&existingUser).Error; err == nil {
		u.log.Error("[UserRepository.Create] User with the same email or username already exists", "email", user.Email, "username", user.Username)
		return models.User{}, errors.New("user with the same email or username already exists")
	}

	if err := u.db.Create(&user).Error; err != nil {
		u.log.Error("[UserRepository.Create] Error while creating user", "error", err.Error())
		return models.User{}, err
	}

	u.log.Info("[UserRepository.Create] User created successfully", "userId", user.ID)
	return user, nil
}

func (u *userRepository) Delete(userId uint) error {
	if err := u.db.Where("id = ?", userId).Delete(models.User{}).Error; err != nil {
		u.log.Error("[UserRepository.Delete] Error while deleting user", "userId", userId, "error", err.Error())
		return err
	}

	u.log.Info("[UserRepository.Delete] User deleted successfully", "userId", userId)
	return nil
}

func (u *userRepository) GetByEmail(email string) (models.User, error) {
	user := models.User{}
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		u.log.Error("[UserRepository.GetByEmail] Error while retrieving user by email", "email", email, "error", err.Error())
		return models.User{}, err
	}

	u.log.Info("[UserRepository.GetByEmail] User retrieved successfully", "userId", user.ID, "email", email)
	return user, nil
}

func (u *userRepository) GetById(userId uint) (models.User, error) {
	user := models.User{}
	if err := u.db.First(&user, userId).Error; err != nil {
		u.log.Error("[UserRepository.GetById] Error while retrieving user by ID", "userId", userId, "error", err.Error())
		return models.User{}, err
	}

	u.log.Info("[UserRepository.GetById] User retrieved successfully", "userId", userId)
	return user, nil
}

func (u *userRepository) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		u.log.Error("[UserRepository.GetByUsername] Error while retrieving user by username", "username", username, "error", err.Error())
		return models.User{}, err
	}

	u.log.Info("[UserRepository.GetByUsername] User retrieved successfully", "userId", user.ID, "username", username)
	return user, nil
}

func (u *userRepository) Update(userId uint, userDTO UpdateUserDTO) (models.User, error) {
	foundUser := models.User{}

	if err := u.db.Where("id = ?", userId).First(&foundUser).Error; err != nil {
		u.log.Error("[UserRepository.Update] Error while retrieving user by ID", "userId", userId, "error", err.Error())
		return models.User{}, err
	}

	updates := map[string]interface{}{}

	if userDTO.Email != nil && *userDTO.Email != "" {
		updates["email"] = *userDTO.Email
	}
	if userDTO.Password != "" {
		updates["password"] = userDTO.Password
	}
	if userDTO.Username != "" {
		updates["username"] = userDTO.Username
	}

	if err := u.db.Model(&foundUser).Updates(updates).Error; err != nil {
		u.log.Error("[UserRepository.Update] Error while updating user", "userId", userId, "error", err.Error())
		return models.User{}, err
	}

	if err := u.db.First(&foundUser, userId).Error; err != nil {
		u.log.Error("[UserRepository.Update] Error while retrieving updated user", "userId", userId, "error", err.Error())
		return models.User{}, err
	}

	u.log.Info("[UserRepository.Update] User updated successfully", "userId", userId)
	return foundUser, nil
}

func New(db *gorm.DB, log *slog.Logger) UserRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}
