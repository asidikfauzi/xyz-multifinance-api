package transaction

import "asidikfauzi/xyz-multifinance-api/internal/model"

type TransactionsMySQL interface {
	FindByContractNumber(string) (model.Transactions, error)
	Transaction(model.Transactions, model.Limits) (model.Transactions, error)
}
