package role

import (
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"github.com/google/uuid"
)

type RolesMySQL interface {
	FindById(uuid uuid.UUID) (model.Roles, error)
}
