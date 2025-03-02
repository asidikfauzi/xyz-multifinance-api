package dto

import (
	"github.com/google/uuid"
)

type (
	RegisterInput struct {
		Email           string    `json:"email" validate:"required,email"`
		Password        string    `json:"password" validate:"required,min=8,password"`
		PasswordConfirm string    `json:"password_confirm" validate:"required,eqfield=Password"`
		RoleId          uuid.UUID `json:"role_id" validate:"required"`
	}

	RegisterResponse struct {
		ID         uuid.UUID `json:"id"`
		Email      string    `json:"email"`
		Role       string    `json:"role"`
		ConsumerId uuid.UUID `json:"consumer_id"`
		CreatedAt  string    `json:"created_at"`
	}
)
