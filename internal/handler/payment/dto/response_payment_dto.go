package dto

import (
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"github.com/google/uuid"
)

type PaymentsResponseWithPage struct {
	Data []PaymentResponse           `json:"data"`
	Page response.PaginationResponse `json:"page"`
}

type PaymentResponse struct {
	ID             uuid.UUID `json:"id"`
	ContractNumber string    `json:"contract_number"`
	Date           *string   `json:"date"`
	AmountPaid     float64   `json:"amount_paid"`
	Status         string    `json:"status"`
	CreatedAt      *string   `json:"created_at,omitempty"`
}
