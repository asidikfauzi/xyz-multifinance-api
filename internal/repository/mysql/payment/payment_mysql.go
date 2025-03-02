package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"github.com/google/uuid"
)

type PaymentsMySQL interface {
	CountPaymentsByCustomerID(uuid.UUID) (int64, error)
	Create(*model.Payments) (model.Payments, error)
}
