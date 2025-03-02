package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"github.com/google/uuid"
)

type PaymentsMySQL interface {
	FindById(uuid uuid.UUID) (model.Payments, error)
}
