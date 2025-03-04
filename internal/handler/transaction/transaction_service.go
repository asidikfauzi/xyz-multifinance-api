package transaction

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/transaction/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/consumer"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/transaction"
	"errors"
	"github.com/google/uuid"
	"math"
	"net/http"
	"time"
)

type transactionService struct {
	transactionMySQL transaction.TransactionsMySQL
	consumerMySQL    consumer.ConsumersMySQL
}

func NewTransactionService(tt transaction.TransactionsMySQL, cm consumer.ConsumersMySQL) TransactionsService {
	return &transactionService{
		transactionMySQL: tt,
		consumerMySQL:    cm,
	}
}

func normalizeTransactionQuery(q dto.QueryTransaction) dto.QueryTransaction {
	if q.Page == 0 {
		q.Page = 1
	}
	if q.Limit == 0 {
		q.Limit = 10
	}
	return q
}

func (s *transactionService) FindAll(q dto.QueryTransaction) (res dto.TransactionsResponseWithPage, code int, err error) {
	q = normalizeTransactionQuery(q)

	transactionData, totalItems, err := s.transactionMySQL.FindAll(q)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	for _, c := range transactionData {
		res.Data = append(res.Data, dto.TransactionsResponse{
			ID:             c.ID,
			NIK:            *c.Consumer.NIK,
			FullName:       c.Consumer.FullName,
			ContractNumber: c.ContractNumber,
			OTR:            c.OTR,
			Tenor:          c.Tenor,
			AdminFee:       c.AdminFee,
			InstallmentAmt: c.InstallmentAmt,
			AmountInterest: c.AmountInterest,
			AssetName:      c.AssetName,
		})
	}

	res.Page = response.PaginationResponse{
		TotalItems:   totalItems,
		ItemCount:    len(res.Data),
		ItemsPerPage: q.Limit,
		CurrentPage:  q.Page,
	}

	return res, http.StatusOK, nil
}

func (s *transactionService) Transaction(input dto.TransactionInput) (res dto.TransactionsResponse, code int, err error) {
	checkConsumer, err := s.consumerMySQL.FindById(input.ConsumerID)
	if err != nil {
		if errors.Is(err, constant.ConsumerNotFound) {
			return res, http.StatusNotFound, constant.ConsumerNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	if !checkConsumer.IsVerified {
		return res, http.StatusForbidden, constant.ConsumerHasNoLimit
	}

	r := (float64(constant.Interest) / 100) / 12
	n := float64(input.Tenor)
	P := input.OTR

	totalAdminFee := constant.AdminFee * n

	//Cicilan
	installmentRaw := (P * r) / (1 - math.Pow(1+r, -n))
	installment := math.Round(installmentRaw*100) / 100

	//Bunga
	amountInterest := math.Round((installment*n-P)*100) / 100

	// Cicilan + Bunga
	installment = installment + constant.AdminFee

	// Total
	total := installment * n

	if total > checkConsumer.Limits[0].LimitAvailable {
		return res, http.StatusForbidden, constant.InsufficientLimit
	}

	contractNumber := utils.GenerateContractNumber()

	createTransaction := model.Transactions{
		ID:             uuid.New(),
		ContractNumber: contractNumber,
		OTR:            math.Round(P*100) / 100,
		Tenor:          input.Tenor,
		AdminFee:       totalAdminFee,
		InstallmentAmt: installment,
		AmountInterest: amountInterest,
		AssetName:      input.AssetName,
		ConsumerID:     checkConsumer.ID,
		CreatedAt:      time.Now(),
		CreatedBy:      input.CreatedBy,
	}

	newLimitAvailable := checkConsumer.Limits[0].LimitAvailable - total

	updateLimit := model.Limits{
		ID:             checkConsumer.Limits[0].ID,
		LimitAvailable: math.Round(newLimitAvailable*100) / 100,
		UpdatedBy:      &input.ConsumerID,
	}

	newTransaction, err := s.transactionMySQL.Transaction(createTransaction, updateLimit)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = dto.TransactionsResponse{
		ID:             newTransaction.ID,
		NIK:            *checkConsumer.NIK,
		FullName:       checkConsumer.FullName,
		LimitAvailable: updateLimit.LimitAvailable,
		ContractNumber: newTransaction.ContractNumber,
		OTR:            newTransaction.OTR,
		Tenor:          newTransaction.Tenor,
		AdminFee:       newTransaction.AdminFee,
		InstallmentAmt: newTransaction.InstallmentAmt,
		AmountInterest: newTransaction.AmountInterest,
		AssetName:      newTransaction.AssetName,
		CreatedAt:      utils.FormatTime(newTransaction.CreatedAt),
	}

	return res, http.StatusCreated, nil
}
