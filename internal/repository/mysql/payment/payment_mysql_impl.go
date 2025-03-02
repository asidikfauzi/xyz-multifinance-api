package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"errors"
	"github.com/google/uuid"
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

func (r *paymentMysql) CountPaymentsByCustomerID(customerID uuid.UUID) (count int64, err error) {
	err = r.DB.Table("payments").
		Select("COUNT(payments.id)").
		Joins("JOIN transactions ON transactions.id = payments.transaction_id").
		Where("transactions.consumer_id = ? AND payments.deleted_at IS NULL", customerID).
		Count(&count).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, constant.CountPaymentNotFound
	}

	return count, err
}

func (r *paymentMysql) Create(input *model.Payments) (res model.Payments, err error) {
	err = r.DB.Create(input).Error
	if err != nil {
		return res, err
	}

	return *input, nil
}
