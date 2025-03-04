package mocks

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/consumer/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type ConsumerMySQLRepository struct {
	mock.Mock
}

func (m *ConsumerMySQLRepository) FindAll(q dto.QueryConsumer) ([]model.Consumers, int64, error) {
	args := m.Called(q)
	if args.Get(0) != nil {
		return args.Get(0).([]model.Consumers), args.Get(1).(int64), args.Error(2)
	}
	return nil, 0, args.Error(2)
}

func (m *ConsumerMySQLRepository) FindById(id uuid.UUID) (model.Consumers, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(model.Consumers), args.Error(1)
	}
	return model.Consumers{}, constant.ConsumerNotFound
}

func (m *ConsumerMySQLRepository) FindByNIK(nik string) (model.Consumers, error) {
	args := m.Called(nik)
	if args.Get(0) != nil {
		return args.Get(0).(model.Consumers), args.Error(1)
	}
	return model.Consumers{}, constant.ConsumerNotFound
}

func (m *ConsumerMySQLRepository) Update(input model.Consumers) (model.Consumers, error) {
	args := m.Called(input)
	if args.Get(0) != nil {
		return args.Get(0).(model.Consumers), args.Error(1)
	}
	return model.Consumers{}, args.Error(1)
}
