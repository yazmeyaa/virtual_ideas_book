// Authentication request DTOs
package request

type RegisterUserDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"omitempty,email"`
}
