package transaction

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/transaction/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"errors"
	"fmt"
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

func (t *transactionMySQL) FindAll(q dto.QueryTransaction) (res []model.Transactions, totalItem int64, err error) {
	if q.OrderBy == "" {
		q.OrderBy = "created_at"
	}

	if q.Direction != "ASC" && q.Direction != "DESC" {
		q.Direction = "DESC"
	}

	offset := (q.Page - 1) * q.Limit

	query := t.DB.Model(&model.Transactions{}).Preload("Consumer").Where("deleted_at IS NULL")

	if q.Search != "" {
		query = query.Where("contract_number LIKE ? OR asset_name LIKE ? OR consumers.full_name LIKE ? OR consumers.nik LIKE ?", "%"+q.Search+"%", "%"+q.Search+"%", "%"+q.Search+"%", "%"+q.Search+"%")
	}

	if err = query.Count(&totalItem).Error; err != nil {
		return nil, 0, err
	}

	if q.ConsumerId != "" {
		query = query.Where("consumer_id = ?", q.ConsumerId)
	}

	if q.Limit > 0 {
		query = query.Limit(q.Limit)
	}

	query = query.Offset(offset)

	if err = query.Order(fmt.Sprintf("%s %s", q.OrderBy, q.Direction)).Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, totalItem, nil
}

func (t *transactionMySQL) FindByContractNumber(contractNumber string) (res model.Transactions, err error) {
	err = t.DB.
		Preload("Consumer").
		Preload("Consumer.Limits").
		Where("deleted_at IS NULL AND contract_number = ?", contractNumber).
		First(&res).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, constant.ContractNumberNotFound
	}

	return res, err
}

func (t *transactionMySQL) Transaction(transactions model.Transactions, limits model.Limits) (res model.Transactions, err error) {
	tx := t.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err = tx.Create(&transactions).Error; err != nil {
		return res, err
	}

	if err = tx.Model(&model.Limits{}).Where("id = ?", limits.ID).
		Update("limit_available", limits.LimitAvailable).Error; err != nil {
		return res, err
	}

	return transactions, nil
}
