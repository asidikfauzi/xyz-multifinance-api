package consumer

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/consumer/dto"
	"github.com/google/uuid"
)

type ConsumersService interface {
	FindAll(dto.QueryConsumer) (dto.ConsumersResponseWithPage, int, error)
	FindById(uuid.UUID) (dto.ConsumerResponse, int, error)
	Create(dto.CreateConsumerInput) (dto.ConsumerResponse, int, error)
	Update(uuid.UUID, dto.UpdateConsumerInput) (dto.ConsumerResponse, int, error)
}
