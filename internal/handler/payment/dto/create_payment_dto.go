package dto

type PaymentInput struct {
	AmountPaid     float64 `json:"amount_paid"`
	ContractNumber string  `json:"contract_number"`
}
