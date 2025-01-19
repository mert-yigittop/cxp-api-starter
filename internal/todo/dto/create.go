package dto

import "github.com/invopop/validation"

type CreateTodoRequest struct {
	Content string `json:"content"`
}

type CreateTodoResponse struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

func (req CreateTodoRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Content, validation.Required, validation.Length(1, 255)),
	)
}
