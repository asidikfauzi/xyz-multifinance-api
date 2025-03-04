package transaction

import "asidikfauzi/xyz-multifinance-api/internal/handler/transaction/dto"

type TransactionsService interface {
	FindAll(dto.QueryTransaction) (dto.TransactionsResponseWithPage, int, error)
	Transaction(dto.TransactionInput) (dto.TransactionsResponse, int, error)
}
