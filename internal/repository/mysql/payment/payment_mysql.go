package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"github.com/google/uuid"
)

type PaymentsMySQL interface {
	CountPaymentsByConsumerID(uuid.UUID) (int64, error)
	Create(*model.Payments, *model.Limits) (model.Payments, error)
}
