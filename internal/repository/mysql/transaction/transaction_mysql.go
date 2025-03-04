package transaction

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/transaction/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
)

type TransactionsMySQL interface {
	FindAll(dto.QueryTransaction) ([]model.Transactions, int64, error)
	FindByContractNumber(string) (model.Transactions, error)
	Transaction(model.Transactions, model.Limits) (model.Transactions, error)
}
