package payment

import "asidikfauzi/xyz-multifinance-api/internal/handler/payment/dto"

type PaymentsService interface {
	Create(dto.PaymentInput) (dto.PaymentResponse, int, error)
}
