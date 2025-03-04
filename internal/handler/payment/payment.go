package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/payment/dto"
	"github.com/google/uuid"
)

type PaymentsService interface {
	FindAll(dto.QueryPayment) (dto.PaymentsResponseWithPage, int, error)
	Pay(uuid.UUID, dto.PaymentInput) (dto.PaymentResponse, int, error)
}
