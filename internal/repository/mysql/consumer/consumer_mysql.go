package consumer

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/consumer/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"github.com/google/uuid"
)

type ConsumersMySQL interface {
	FindAll(dto.QueryConsumer) ([]model.Consumers, int64, error)
	FindById(uuid.UUID) (model.Consumers, error)
	FindByNIK(string) (model.Consumers, error)
	Update(model.Consumers) (model.Consumers, error)
}
