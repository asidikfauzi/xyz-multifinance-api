package service

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/transaction"
	"asidikfauzi/xyz-multifinance-api/internal/handler/transaction/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/test/mocks"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math"
	"net/http"
	"testing"
)

func TestTransaction_Success(t *testing.T) {
	mockConsumerRepo := new(mocks.ConsumerMySQLRepository)
	mockTransactionRepo := new(mocks.TransactionMySQLRepository)

	transactionService := transaction.NewTransactionService(mockTransactionRepo, mockConsumerRepo)

	consumerID := uuid.New()
	transactionInput := dto.TransactionInput{
		ConsumerID: consumerID,
		OTR:        1000000,
		Tenor:      60,
		AssetName:  "Motor",
		CreatedBy:  uuid.New(),
	}

	nik := "1234567890123456"
	mockConsumer := model.Consumers{
		ID:              consumerID,
		NIK:             &nik,
		FullName:        "John Doe",
		LegalName:       "John",
		Phone:           "087890987890",
		PlaceOfBirth:    "Sumenep",
		DateOfBirth:     "01-01-2001",
		Salary:          10000000,
		KTPImage:        "image.png",
		SelfieImage:     "image.png",
		IsVerified:      true,
		RejectionReason: "",
		Limits: []model.Limits{
			{ID: uuid.New(), LimitAvailable: 20000000},
		},
	}

	r := (float64(constant.Interest) / 100) / 12
	n := float64(transactionInput.Tenor)
	P := transactionInput.OTR

	totalAdminFee := constant.AdminFee * n

	//Cicilan
	installmentRaw := (P * r) / (1 - math.Pow(1+r, -n))
	installment := math.Round(installmentRaw*100) / 100

	//Bunga
	amountInterest := math.Round((installment*n-P)*100) / 100

	// Cicilan + Bunga
	installment = installment + constant.AdminFee

	// Total
	total := installment*n + totalAdminFee

	fmt.Println("OTR:", transactionInput.OTR)
	fmt.Println("Tenor:", transactionInput.Tenor)
	fmt.Println("Admin Fee:", constant.AdminFee)
	fmt.Println("Installment:", installment)
	fmt.Println("amountInterest:", amountInterest)
	fmt.Println("Total:", total)
	fmt.Println("Consumer Limit Available:", mockConsumer.Limits[0].LimitAvailable)

	if total > mockConsumer.Limits[0].LimitAvailable {
		t.Fatalf("Limit Tidak Cukup! Seharusnya test mengembalikan 403")
	}

	mockTransaction := model.Transactions{
		ID:             uuid.New(),
		ContractNumber: "ABC123",
		OTR:            transactionInput.OTR,
		Tenor:          transactionInput.Tenor,
		AdminFee:       constant.AdminFee,
		InstallmentAmt: installment,
		AmountInterest: amountInterest,
		AssetName:      transactionInput.AssetName,
		ConsumerID:     consumerID,
	}

	mockConsumerRepo.On("FindById", transactionInput.ConsumerID).Return(mockConsumer, nil)
	mockTransactionRepo.On("Transaction", mock.Anything, mock.Anything).Return(mockTransaction, nil)

	res, code, err := transactionService.Transaction(transactionInput)

	fmt.Println("Response:", res)
	fmt.Println("HTTP Code:", code)
	fmt.Println("Error:", err)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, code)
	assert.Equal(t, mockTransaction.ContractNumber, res.ContractNumber)
	mockConsumerRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}

func TestTransaction_ConsumerNotFound(t *testing.T) {
	mockConsumerRepo := new(mocks.ConsumerMySQLRepository)
	mockTransactionRepo := new(mocks.TransactionMySQLRepository)

	transactionService := transaction.NewTransactionService(mockTransactionRepo, mockConsumerRepo)

	consumerID := uuid.New()
	transactionInput := dto.TransactionInput{
		ConsumerID: consumerID,
		OTR:        10000000,
		Tenor:      12,
		AssetName:  "Motor",
		CreatedBy:  uuid.New(),
	}

	mockConsumerRepo.On("FindById", consumerID).Return(model.Consumers{}, constant.ConsumerNotFound)

	_, code, err := transactionService.Transaction(transactionInput)

	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, code)
	assert.Equal(t, constant.ConsumerNotFound, err)
	mockConsumerRepo.AssertExpectations(t)
}

func TestTransaction_InsufficientLimit(t *testing.T) {
	mockConsumerRepo := new(mocks.ConsumerMySQLRepository)
	mockTransactionRepo := new(mocks.TransactionMySQLRepository)

	transactionService := transaction.NewTransactionService(mockTransactionRepo, mockConsumerRepo)

	consumerID := uuid.New()
	transactionInput := dto.TransactionInput{
		ConsumerID: consumerID,
		OTR:        20000000,
		Tenor:      12,
		AssetName:  "Mobil",
		CreatedBy:  uuid.New(),
	}

	mockConsumer := model.Consumers{
		ID:         consumerID,
		FullName:   "John Doe",
		IsVerified: true,
		Limits: []model.Limits{
			{ID: uuid.New(), LimitAvailable: 10000000},
		},
	}

	mockConsumerRepo.On("FindById", transactionInput.ConsumerID).Return(mockConsumer, nil)

	_, code, err := transactionService.Transaction(transactionInput)

	assert.Error(t, err)
	assert.Equal(t, http.StatusForbidden, code)
	assert.Equal(t, constant.InsufficientLimit, err)
	mockConsumerRepo.AssertExpectations(t)
}
