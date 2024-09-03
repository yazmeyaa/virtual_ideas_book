package request

type LoginUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email,omitempty"`
}
