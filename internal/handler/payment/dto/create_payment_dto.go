package dto

import "github.com/google/uuid"

type PaymentInput struct {
	AmountPaid     float64   `json:"amount_paid" validate:"required,numeric"`
	ContractNumber string    `json:"contract_number" validate:"required"`
	ConsumerId     uuid.UUID `json:"consumer_id" validate:"required,uuid"`
	UpdatedBy      uuid.UUID `json:"updated_by" validate:"required,uuid"`
}
