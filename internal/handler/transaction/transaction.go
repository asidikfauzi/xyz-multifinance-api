package transaction

import "asidikfauzi/xyz-multifinance-api/internal/handler/transaction/dto"

type TransactionsService interface {
	Transaction(dto.TransactionInput) (dto.TransactionsResponse, int, error)
}
