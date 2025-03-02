package role

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type roleMysql struct {
	DB *gorm.DB
}

func NewRolesMySQL(db *gorm.DB) RolesMySQL {
	return &roleMysql{
		DB: db,
	}
}

func (r *roleMysql) FindById(id uuid.UUID) (res model.Roles, err error) {
	err = r.DB.Where("id = ?", id).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, constant.RoleNotFound
	}

	return res, err
}
