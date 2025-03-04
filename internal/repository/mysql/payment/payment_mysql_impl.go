package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"errors"
	"gorm.io/gorm"
)

type paymentMysql struct {
	DB *gorm.DB
}

func NewPaymentsMySQL(db *gorm.DB) PaymentsMySQL {
	return &paymentMysql{
		DB: db,
	}
}

func (p *paymentMysql) CountPaymentsByContractNumber(contractNumber string) (count int64, err error) {
	err = p.DB.Table("payments").
		Select("COUNT(payments.id)").
		Joins("JOIN transactions ON transactions.id = payments.transaction_id").
		Where("transactions.contract_number = ? AND payments.deleted_at IS NULL", contractNumber).
		Count(&count).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, constant.CountPaymentNotFound
	}

	return count, err
}

func (p *paymentMysql) Create(payment *model.Payments, limit *model.Limits) (res model.Payments, err error) {
	tx := p.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := tx.Create(payment).Error; err != nil {
		return res, err
	}

	if err := tx.Model(&model.Limits{}).
		Where("id = ?", limit.ID).
		Update("limit_available", limit.LimitAvailable).Error; err != nil {
		return res, err
	}

	return *payment, nil
}
