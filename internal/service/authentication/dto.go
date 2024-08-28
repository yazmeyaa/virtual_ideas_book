package authentication

type RegisterUserDTO struct {
	Username string
	Password string
	Email    string
}

type LoginDTO struct {
	Username string
	Password string
}
