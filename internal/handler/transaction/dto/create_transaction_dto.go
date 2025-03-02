package dto

import "github.com/google/uuid"

type TransactionInput struct {
	OTR        float64   `json:"otr" validate:"required,numeric"`
	AssetName  string    `json:"asset_name" validate:"required"`
	Tenor      int       `json:"tenor" validate:"required,numeric"`
	ConsumerID uuid.UUID `json:"consumer_id" validate:"required,uuid"`
	CreatedBy  uuid.UUID `json:"created_by" validate:"required,uuid"`
}
