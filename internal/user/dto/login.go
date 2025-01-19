package dto

import "github.com/invopop/validation"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID          uint   `json:"id"`
	AccessToken string `json:"accessToken"`
}

func (req LoginRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Username, validation.Required, validation.Length(3, 255)),
		validation.Field(&req.Password, validation.Required, validation.Length(8, 255)),
	)
}
