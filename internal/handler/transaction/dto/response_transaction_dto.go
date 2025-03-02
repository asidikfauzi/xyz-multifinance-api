package dto

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/consumer/dto"
	dto2 "asidikfauzi/xyz-multifinance-api/internal/handler/limit/dto"
	"github.com/google/uuid"
	"time"
)

type TransactionsResponse struct {
	ID             uuid.UUID            `json:"id"`
	ContractNumber string               `json:"contract_number"`
	OTR            float64              `json:"otr"`
	Tenor          int                  `json:"tenor"`
	AdminFee       float64              `json:"admin_fee"`
	InstallmentAmt float64              `json:"installment_amt"`
	AmountInterest float64              `json:"amount_interest"`
	AssetName      string               `json:"asset_name"`
	Consumer       dto.ConsumerResponse `json:"consumer"`
	Limit          dto2.LimitResponse   `json:"limit"`
	CreatedAt      time.Time            `json:"created_at"`
}
