package dto

import "github.com/invopop/validation"

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	ID uint `json:"id"`
}

func (req RegisterRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required, validation.Length(3, 255)),
		validation.Field(&req.Email, validation.Required, validation.Length(8, 255)),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 255)),
	)
}
