package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/payment/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/payment"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/transaction"
	"errors"
	"github.com/google/uuid"
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

func (r *paymentService) Create(input dto.PaymentInput) (res dto.PaymentResponse, code int, err error) {
	transactionData, err := r.transactionMySQL.FindByContractNumber(input.ContractNumber)
	if err != nil {
		if errors.Is(err, constant.ContractNumberNotFound) {
			return res, http.StatusNotFound, constant.ContractNumberNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	if transactionData.ConsumerID != input.ConsumerId {
		return res, http.StatusForbidden, constant.AccessDenied
	}

	if transactionData.InstallmentAmt != input.AmountPaid {
		return res, http.StatusBadRequest, constant.AmountPaidMustBeEqual
	}

	countPayment, err := r.paymentMySQL.CountPaymentsByCustomerID(transactionData.ConsumerID)
	if err != nil {
		if errors.Is(err, constant.CountPaymentNotFound) {
			return res, http.StatusNotFound, constant.CountPaymentNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	if countPayment == int64(transactionData.Tenor) {
		return res, http.StatusConflict, constant.PaymentAlreadyComplete
	}

	paymentData := model.Payments{
		ID:            uuid.New(),
		Date:          time.Now(),
		AmountPaid:    input.AmountPaid,
		Status:        string(constant.SUCCESS),
		TransactionID: transactionData.ID,
		CreatedBy:     input.CreatedBy,
	}

	newPayment, err := r.paymentMySQL.Create(&paymentData)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = dto.PaymentResponse{
		ID:         newPayment.ID,
		Date:       newPayment.Date,
		AmountPaid: newPayment.AmountPaid,
		Status:     newPayment.Status,
		CreatedAt:  newPayment.CreatedAt,
	}

	return res, http.StatusCreated, nil
}
