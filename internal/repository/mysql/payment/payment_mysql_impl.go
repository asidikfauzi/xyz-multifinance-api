package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/payment/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"errors"
	"fmt"
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

func (p *paymentMysql) FindById(id uuid.UUID) (res model.Payments, err error) {
	err = p.DB.Where("deleted_at IS NULL AND id = ?", id).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, constant.PaymentNotFound
	}

	return res, err
}

func (p *paymentMysql) FindAll(q dto.QueryPayment) (res []model.Payments, totalItem int64, err error) {
	if q.OrderBy == "" {
		q.OrderBy = "created_at"
	}

	if q.Direction != "ASC" && q.Direction != "DESC" {
		q.Direction = "DESC"
	}

	offset := (q.Page - 1) * q.Limit

	query := p.DB.Model(&model.Payments{}).
		Joins("JOIN transactions ON transactions.id = payments.transaction_id").
		Preload("Transaction").
		Where("payments.deleted_at IS NULL")

	if q.Search != "" {
		query = query.Where("transactions.contract_number LIKE ? OR status LIKE ? OR transactions.consumer_id LIKE ?", "%"+q.Search+"%", "%"+q.Search+"%", "%"+q.Search+"%")
	}

	if q.Status != "" {
		query = query.Where("status = ?", q.Status)
	}

	if q.ConsumerId != "" {
		query = query.Where("transactions.consumer_id = ?", q.ConsumerId)
	}

	if q.ContractNumber != "" {
		query = query.Where("transactions.contract_number = ?", q.ContractNumber)
	}

	if q.Limit > 0 {
		query = query.Limit(q.Limit)
	}

	if err = query.Count(&totalItem).Error; err != nil {
		return nil, 0, err
	}

	query = query.Offset(offset)

	if err = query.Order(fmt.Sprintf("%s %s", q.OrderBy, q.Direction)).Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, totalItem, nil
}

func (p *paymentMysql) Pay(payment *model.Payments, limit *model.Limits) (res model.Payments, err error) {
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

	if err := tx.Model(&model.Payments{}).
		Where("id = ?", payment.ID).
		Updates(payment).Error; err != nil {
		return res, err
	}

	if err := tx.Model(&model.Limits{}).
		Where("id = ?", limit.ID).
		Update("limit_available", limit.LimitAvailable).Error; err != nil {
		return res, err
	}

	return *payment, nil
}
