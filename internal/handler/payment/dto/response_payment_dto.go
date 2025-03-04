package dto

import (
	"github.com/google/uuid"
	"time"
)

type PaymentResponse struct {
	ID         uuid.UUID `json:"id"`
	Date       time.Time `json:"date"`
	AmountPaid float64   `json:"amount_paid"`
	Status     string    `json:"status"`
	CreatedAt  *string   `json:"created_at"`
}
