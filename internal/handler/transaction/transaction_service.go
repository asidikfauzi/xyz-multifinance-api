package transaction

import (
	dto2 "asidikfauzi/xyz-multifinance-api/internal/handler/consumer/dto"
	dto3 "asidikfauzi/xyz-multifinance-api/internal/handler/limit/dto"
	"asidikfauzi/xyz-multifinance-api/internal/handler/transaction/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/consumer"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/transaction"
	"errors"
	"github.com/google/uuid"
	"math"
	"net/http"
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

	// Perhitungan Total Pinjaman & Angsuran
	r := (float64(constant.Interest) / 100) / 12 // Bunga per bulan
	n := float64(input.Tenor)
	P := input.OTR // Pokok pinjaman

	// Rumus annuitas
	installmentRaw := (P * r) / (1 - math.Pow(1+r, -n))
	installment := math.Round(installmentRaw*100) / 100 // Bulatkan setelah semua perhitungan

	// Hitung total pembayaran
	total := P + (installment * n) + float64(constant.AdminFee)

	// Hitung total bunga setelah installment dibulatkan
	amountInterest := math.Round((installment*n-P)*100) / 100

	// Cek apakah limit cukup sebelum lanjut transaksi
	if total > checkConsumer.Limits[0].LimitAvailable {
		return res, http.StatusForbidden, constant.InsufficientLimit
	}

	contractNumber := utils.GenerateContractNumber()

	createTransaction := model.Transactions{
		ID:             uuid.New(),
		ContractNumber: contractNumber,
		OTR:            math.Round(P*100) / 100,
		Tenor:          input.Tenor,
		AdminFee:       constant.AdminFee,
		InstallmentAmt: installment,
		AmountInterest: amountInterest,
		AssetName:      input.AssetName,
		ConsumerID:     checkConsumer.ID,
		CreatedBy:      input.CreatedBy,
	}

	newLimitAvailable := checkConsumer.Limits[0].LimitAvailable - total

	updateLimit := model.Limits{
		ID:             checkConsumer.Limits[0].ID,
		LimitAvailable: math.Round(newLimitAvailable*100) / 100,
	}

	newTransaction, err := s.transactionMySQL.Transaction(createTransaction, updateLimit)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	consumerData := dto2.ConsumerResponse{
		ID:              checkConsumer.ID,
		NIK:             *checkConsumer.NIK,
		FullName:        checkConsumer.FullName,
		LegalName:       checkConsumer.LegalName,
		Phone:           checkConsumer.Phone,
		PlaceOfBirth:    checkConsumer.PlaceOfBirth,
		DateOfBirth:     checkConsumer.DateOfBirth,
		Salary:          checkConsumer.Salary,
		KtpImage:        checkConsumer.KTPImage,
		SelfieImage:     checkConsumer.SelfieImage,
		IsVerified:      checkConsumer.IsVerified,
		RejectionReason: checkConsumer.RejectionReason,
	}

	limitData := dto3.LimitResponse{
		ID:             updateLimit.ID,
		LimitAvailable: updateLimit.LimitAvailable,
	}

	res = dto.TransactionsResponse{
		ID:             newTransaction.ID,
		ContractNumber: newTransaction.ContractNumber,
		OTR:            newTransaction.OTR,
		Tenor:          newTransaction.Tenor,
		AdminFee:       newTransaction.AdminFee,
		InstallmentAmt: newTransaction.InstallmentAmt,
		AmountInterest: newTransaction.AmountInterest,
		AssetName:      newTransaction.AssetName,
		Consumer:       consumerData,
		Limit:          limitData,
	}

	return res, http.StatusCreated, nil
}
