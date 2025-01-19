package dto

import "github.com/invopop/validation"

type UpdateTodoRequest struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"userId"`
}

type UpdateTodoResponse struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

func (req UpdateTodoRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Content, validation.Length(1, 255)),
		validation.Field(&req.Completed),
	)
}
