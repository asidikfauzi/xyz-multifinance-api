package consumer

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/consumer/dto"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type consumerMysql struct {
	DB *gorm.DB
}

func NewConsumersMySQL(db *gorm.DB) ConsumersMySQL {
	return &consumerMysql{
		DB: db,
	}
}

func (c *consumerMysql) FindAll(q dto.QueryConsumer) (res []model.Consumers, totalItem int64, err error) {
	if q.OrderBy == "" {
		q.OrderBy = "created_at"
	}

	if q.Direction != "ASC" && q.Direction != "DESC" {
		q.Direction = "DESC"
	}

	offset := (q.Page - 1) * q.Limit

	query := c.DB.Model(&model.Consumers{}).Where("deleted_at IS NULL")

	if q.Search != "" {
		query = query.Where("nik LIKE ? OR full_name LIKE ? OR legal_name LIKE ? OR place_of_birth LIKE ?", "%"+q.Search+"%", "%"+q.Search+"%", "%"+q.Search+"%", "%"+q.Search+"%")
	}

	if err = query.Count(&totalItem).Error; err != nil {
		return nil, 0, err
	}

	if q.Paginate == "true" {
		if q.Limit > 0 {
			query = query.Limit(q.Limit)
		}

		query = query.Offset(offset)
	}

	if err = query.Order(fmt.Sprintf("%s %s", q.OrderBy, q.Direction)).Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, totalItem, nil
}

func (c *consumerMysql) FindById(id uuid.UUID) (res model.Consumers, err error) {
	err = c.DB.Preload("Limits").Where("id = ?", id).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, constant.ConsumerNotFound
	}

	return res, err
}

func (c *consumerMysql) FindByNIK(id string) (res model.Consumers, err error) {
	err = c.DB.Where("nik = ?", id).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, constant.ConsumerNotFound
	}

	return res, err
}

func (c *consumerMysql) Create(input model.Consumers) (res model.Consumers, err error) {
	if err = c.DB.Create(&input).Error; err != nil {
		return res, err
	}

	return input, nil
}

func (c *consumerMysql) Update(input model.Consumers) (res model.Consumers, err error) {
	if err = c.DB.Updates(&input).Error; err != nil {
		return res, err
	}

	return input, nil
}
