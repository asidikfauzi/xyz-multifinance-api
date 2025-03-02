package transaction

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"errors"
	"gorm.io/gorm"
)

type transactionMySQL struct {
	DB *gorm.DB
}

func NewTransactionsMySQL(db *gorm.DB) TransactionsMySQL {
	return &transactionMySQL{
		DB: db,
	}
}

func (t *transactionMySQL) FindByContractNumber(contractNumber string) (res model.Transactions, err error) {
	err = t.DB.Where("deleted_at IS NULL AND contract_number = ?", contractNumber).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, constant.ContractNumberNotFound
	}

	return res, err
}

func (t *transactionMySQL) Transaction(transactions model.Transactions, limits model.Limits) (res model.Transactions, err error) {
	tx := t.DB.Begin()

	if err = t.DB.Create(&transactions).Error; err != nil {
		return res, err
	}

	if err = t.DB.Updates(&limits).Error; err != nil {
		return res, err
	}

	tx.Commit()

	return transactions, nil
}
