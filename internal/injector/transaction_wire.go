//go:build wireinject
// +build wireinject

package injector

import (
	"asidikfauzi/xyz-multifinance-api/internal/database"
	"asidikfauzi/xyz-multifinance-api/internal/handler/transaction"
	consumerMySQL "asidikfauzi/xyz-multifinance-api/internal/repository/mysql/consumer"
	transactionMySQL "asidikfauzi/xyz-multifinance-api/internal/repository/mysql/transaction"
	"github.com/google/wire"
)

func InitializedTransactionModule() *transaction.TransactionsController {
	wire.Build(
		database.InitDatabase,
		transactionMySQL.NewTransactionsMySQL,
		consumerMySQL.NewConsumersMySQL,
		transaction.NewTransactionService,
		transaction.NewTransactionsController,
	)

	return nil
}
