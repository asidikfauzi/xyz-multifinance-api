package dto

import (
	"github.com/google/uuid"
)

type LimitResponse struct {
	ID             uuid.UUID `json:"id"`
	Tenor          int       `json:"tenor"`
	LimitAvailable float64   `json:"limit_available"`
	CreatedAt      *string   `json:"created_at,omitempty"`
	UpdatedAt      *string   `json:"updated_at,omitempty"`
}
