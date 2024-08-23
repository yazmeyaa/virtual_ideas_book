package user

type CreateUserDTO struct {
	Username string
	Password string
	Email    *string
}

type UpdateUserDTO struct {
	Username string
	Password string
	Email    *string
}
