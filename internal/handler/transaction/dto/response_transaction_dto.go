package dto

import (
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"github.com/google/uuid"
)

type ConsumersResponseWithPage struct {
	Data []TransactionsResponse      `json:"data"`
	Page response.PaginationResponse `json:"page"`
}

type TransactionsResponse struct {
	ID             uuid.UUID `json:"id"`
	NIK            string    `json:"nik"`
	FullName       string    `json:"full_name"`
	LimitAvailable float64   `json:"limit_available"`
	ContractNumber string    `json:"contract_number"`
	OTR            float64   `json:"otr"`
	Tenor          int       `json:"tenor"`
	AdminFee       float64   `json:"admin_fee"`
	InstallmentAmt float64   `json:"installment_amt"`
	AmountInterest float64   `json:"amount_interest"`
	AssetName      string    `json:"asset_name"`
	CreatedAt      *string   `json:"created_at,omitempty"`
}
