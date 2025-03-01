package auth

import "asidikfauzi/xyz-multifinance-api/internal/handler/auth/dto"

type AuthsService interface {
	Login(dto.LoginInput) (dto.LoginResponse, int, error)
	Register(dto.RegisterInput) (dto.RegisterResponse, int, error)
}
