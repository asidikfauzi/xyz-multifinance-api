package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
)

type PaymentsMySQL interface {
	CountPaymentsByContractNumber(string) (int64, error)
	Create(*model.Payments, *model.Limits) (model.Payments, error)
}
