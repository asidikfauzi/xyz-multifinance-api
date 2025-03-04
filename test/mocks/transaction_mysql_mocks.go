package mocks

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/transaction/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"github.com/stretchr/testify/mock"
)

type TransactionMySQLRepository struct {
	mock.Mock
}

func (m *TransactionMySQLRepository) FindAll(q dto.QueryTransaction) ([]model.Transactions, int64, error) {
	args := m.Called(q)
	if args.Get(0) != nil {
		return args.Get(0).([]model.Transactions), args.Get(1).(int64), args.Error(2)
	}
	return nil, 0, args.Error(2)
}

func (m *TransactionMySQLRepository) FindByContractNumber(contractNumber string) (model.Transactions, error) {
	args := m.Called(contractNumber)
	if args.Get(0) != nil {
		return args.Get(0).(model.Transactions), args.Error(1)
	}
	return model.Transactions{}, args.Error(1)
}

func (m *TransactionMySQLRepository) Transaction(tx model.Transactions, limit model.Limits) (model.Transactions, error) {
	args := m.Called(tx, limit)
	if args.Get(0) != nil {
		return args.Get(0).(model.Transactions), args.Error(1)
	}
	return model.Transactions{}, args.Error(1)
}
