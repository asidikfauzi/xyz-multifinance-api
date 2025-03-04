package mocks

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"github.com/stretchr/testify/mock"
)

type TransactionMySQLRepository struct {
	mock.Mock
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
