package user

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"errors"
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
	if err = r.DB.Create(&user).Error; err != nil {
		return res, err
	}

	if err = r.DB.Preload("Role").First(&res, "id = ?", user.ID).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (r *userMysql) FindByEmail(email string) (res model.Users, err error) {
	err = r.DB.Preload("Role").Where("email = ?", email).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, constant.UserNotFound
	}

	return res, err
}
