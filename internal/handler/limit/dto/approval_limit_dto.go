package dto

import "github.com/google/uuid"

type ApprovalLimitInput struct {
	Tenor           int       `json:"tenor" validate:"max=6,min=0,numeric"`
	LimitAvailable  float64   `json:"limit_available" validate:"numeric,min=0"`
	IsVerified      bool      `json:"is_verified"`
	RejectionReason string    `json:"rejection_reason"`
	CreatedBy       uuid.UUID `json:"created_by" validate:"required"`
}
