package user

import "asidikfauzi/xyz-multifinance-api/internal/model"

type UsersMySQL interface {
	Create(model.Users) (model.Users, error)
	FindByEmail(string) (model.Users, error)
}
