package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/payment/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/payment"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/transaction"
	"errors"
	"github.com/google/uuid"
	"math"
	"net/http"
	"time"
)

type paymentService struct {
	paymentMySQL     payment.PaymentsMySQL
	transactionMySQL transaction.TransactionsMySQL
}

func NewPaymentService(pm payment.PaymentsMySQL, tm transaction.TransactionsMySQL) PaymentsService {
	return &paymentService{
		paymentMySQL:     pm,
		transactionMySQL: tm,
	}
}

func normalizePaymentQuery(q dto.QueryPayment) dto.QueryPayment {
	if q.Page == 0 {
		q.Page = 1
	}
	if q.Limit == 0 {
		q.Limit = 10
	}
	return q
}

func (s *paymentService) FindAll(q dto.QueryPayment) (res dto.PaymentsResponseWithPage, code int, err error) {
	q = normalizePaymentQuery(q)

	transactionData, totalItems, err := s.paymentMySQL.FindAll(q)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	for _, c := range transactionData {
		res.Data = append(res.Data, dto.PaymentResponse{
			ID:             c.ID,
			ContractNumber: c.Transaction.ContractNumber,
			Date:           c.Date,
			AmountPaid:     c.AmountPaid,
			Status:         c.Status,
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

func (s *paymentService) Pay(id uuid.UUID, input dto.PaymentInput) (res dto.PaymentResponse, code int, err error) {
	paymentData, err := s.paymentMySQL.FindById(id)
	if err != nil {
		if errors.Is(err, constant.PaymentNotFound) {
			return res, http.StatusNotFound, constant.PaymentNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	transactionData, err := s.transactionMySQL.FindByContractNumber(input.ContractNumber)
	if err != nil {
		if errors.Is(err, constant.ContractNumberNotFound) {
			return res, http.StatusNotFound, constant.ContractNumberNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	if transactionData.ConsumerID != input.ConsumerId {
		return res, http.StatusNotFound, constant.PaymentNotFound
	}

	if transactionData.InstallmentAmt != input.AmountPaid {
		return res, http.StatusBadRequest, constant.AmountPaidMustBeEqual
	}

	if paymentData.Status == string(constant.SUCCESS) {
		return res, http.StatusConflict, constant.PaymentAlreadyComplete
	}

	paymentData.Status = string(constant.SUCCESS)
	paymentData.UpdatedAt = time.Now()
	paymentData.UpdatedBy = &input.UpdatedBy

	updateLimitRaw := transactionData.Consumer.Limits[0].LimitAvailable + paymentData.AmountPaid
	updateLimit := math.Round(updateLimitRaw*100) / 100

	limitsData := model.Limits{
		ID:             transactionData.Consumer.Limits[0].ID,
		LimitAvailable: updateLimit,
	}

	newPayment, err := s.paymentMySQL.Pay(&paymentData, &limitsData)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = dto.PaymentResponse{
		ID:             newPayment.ID,
		Date:           newPayment.Date,
		ContractNumber: transactionData.ContractNumber,
		AmountPaid:     newPayment.AmountPaid,
		Status:         newPayment.Status,
		CreatedAt:      utils.FormatTime(newPayment.CreatedAt),
	}

	return res, http.StatusOK, nil
}
