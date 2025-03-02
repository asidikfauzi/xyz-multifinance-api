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

func NewRolesMySQL(db *gorm.DB) PaymentsMySQL {
	return &paymentMysql{
		DB: db,
	}
}

func (r *paymentMysql) FindById(id uuid.UUID) (res model.Payments, err error) {
	err = r.DB.Where("deleted_at IS NULL AND id = ?", id).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, constant.RoleNotFound
	}

	return res, err
}

func (r *paymentMysql) FindLastPaymentCustomerID(id uuid.UUID) (res model.Payments, err error) {
	err = r.DB.Where("deleted_at IS NULL AND customer_id = ?", id).
		Order("created_at DESC").
		Limit(1).
		First(&res).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, constant.RoleNotFound
	}

	return res, err
}
