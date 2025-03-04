package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/payment/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"github.com/google/uuid"
)

type PaymentsMySQL interface {
	FindAll(dto.QueryPayment) ([]model.Payments, int64, error)
	FindById(uuid.UUID) (model.Payments, error)
	Pay(*model.Payments, *model.Limits) (model.Payments, error)
}
