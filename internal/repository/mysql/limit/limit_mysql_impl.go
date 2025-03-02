package limit

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"gorm.io/gorm"
)

type limitMysql struct {
	DB *gorm.DB
}

func NewLimitsMySQL(db *gorm.DB) LimitsMySQL {
	return &limitMysql{
		DB: db,
	}
}

func (c *limitMysql) ApprovalConsumer(consumer model.Consumers, limits model.Limits) (res model.Limits, err error) {
	tx := c.DB.Begin()

	if err := tx.Create(&limits).Error; err != nil {
		tx.Rollback()
		return res, err
	}

	if err := tx.Updates(consumer).Error; err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()

	return limits, nil
}
