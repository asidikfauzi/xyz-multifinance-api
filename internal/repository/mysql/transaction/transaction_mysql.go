package transaction

import "asidikfauzi/xyz-multifinance-api/internal/model"

type TransactionsMySQL interface {
	Transaction(model.Transactions, model.Limits) (model.Transactions, error)
}
