package user

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userMysql struct {
	DB *gorm.DB
}

func NewUserMySQL(db *gorm.DB) UsersMySQL {
	return &userMysql{
		DB: db,
	}
}

func (r *userMysql) Create(user model.Users) (res model.Users, err error) {
	tx := r.DB.Begin()

	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return res, err
	}

	consumer := model.Consumers{ID: uuid.New(), UserID: user.ID}
	if err = tx.Create(&consumer).Error; err != nil {
		tx.Rollback()
		return res, err
	}

	if err = tx.Preload("Role").Preload("Consumer").First(&res, "id = ?", user.ID).Error; err != nil {
		tx.Rollback()
		return res, err
	}

	tx.Commit()
	return res, nil
}

func (r *userMysql) FindByEmail(email string) (res model.Users, err error) {
	err = r.DB.Preload("Role").
		Preload("Consumer").
		Where("deleted_at IS NULL AND email = ?", email).
		First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, constant.UserNotFound
	}

	return res, err
}
