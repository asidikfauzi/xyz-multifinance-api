package dto

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/limit/dto"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"github.com/google/uuid"
)

type ConsumersResponseWithPage struct {
	Data []ConsumerResponse          `json:"data"`
	Page response.PaginationResponse `json:"page"`
}

type ConsumerResponse struct {
	ID              uuid.UUID          `json:"id"`
	Email           string             `json:"email"`
	NIK             string             `json:"nik"`
	FullName        string             `json:"full_name"`
	LegalName       string             `json:"legal_name"`
	Phone           string             `json:"phone"`
	PlaceOfBirth    string             `json:"place_of_birth"`
	DateOfBirth     string             `json:"date_of_birth"`
	Salary          float64            `json:"salary"`
	KtpImage        string             `json:"ktp_image"`
	SelfieImage     string             `json:"selfie_image"`
	IsVerified      bool               `json:"is_verified"`
	RejectionReason string             `json:"rejection_reason"`
	CreatedAt       *string            `json:"created_at,omitempty"`
	UpdatedAt       *string            `json:"updated_at,omitempty"`
	Limit           *dto.LimitResponse `json:"limit,omitempty"`
}
