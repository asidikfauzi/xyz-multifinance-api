//go:build wireinject
// +build wireinject

package injector

import (
	"asidikfauzi/xyz-multifinance-api/internal/database"
	"asidikfauzi/xyz-multifinance-api/internal/handler/payment"
	paymentSQL "asidikfauzi/xyz-multifinance-api/internal/repository/mysql/payment"
	transactionSQL "asidikfauzi/xyz-multifinance-api/internal/repository/mysql/transaction"
	"github.com/google/wire"
)

func InitializedPaymentModule() *payment.PaymentsController {
	wire.Build(
		database.InitDatabase,
		paymentSQL.NewPaymentsMySQL,
		transactionSQL.NewTransactionsMySQL,
		payment.NewPaymentService,
		payment.NewPaymentsController,
	)

	return nil
}
