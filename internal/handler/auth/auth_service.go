package auth

import (
	"asidikfauzi/xyz-multifinance-api/internal/config"
	"asidikfauzi/xyz-multifinance-api/internal/handler/auth/dto"
	"asidikfauzi/xyz-multifinance-api/internal/middleware"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/role"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/user"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type authService struct {
	userMySQL user.UsersMySQL
	roleMySQL role.RolesMySQL
}

func NewAuthService(u user.UsersMySQL, r role.RolesMySQL) AuthsService {
	return &authService{
		userMySQL: u,
		roleMySQL: r,
	}
}

func (a *authService) Register(registerInput dto.RegisterInput) (res dto.RegisterResponse, code int, err error) {
	if _, err := a.roleMySQL.FindById(registerInput.RoleID); err != nil {
		if errors.Is(err, constant.RoleNotFound) {
			return res, http.StatusNotFound, constant.RoleNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	if _, err := a.userMySQL.FindByEmail(registerInput.Email); err == nil {
		return res, http.StatusConflict, constant.EmailAlreadyExists
	}

	userData := model.Users{
		ID:       uuid.New(),
		Email:    registerInput.Email,
		Password: utils.HashPassword(registerInput.Password),
		RoleID:   registerInput.RoleID,
	}

	createdUser, err := a.userMySQL.Create(userData)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	localTime, err := utils.FormatTimeWithTimezone(createdUser.CreatedAt)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	return dto.RegisterResponse{
		ID:        createdUser.ID,
		Email:     createdUser.Email,
		Role:      createdUser.Role.Name,
		CreatedAt: localTime,
	}, http.StatusCreated, nil
}

func (a *authService) Login(loginInput dto.LoginInput) (res dto.LoginResponse, code int, err error) {
	userData, err := a.userMySQL.FindByEmail(loginInput.Email)
	if err != nil {
		if errors.Is(err, constant.UserNotFound) {
			return res, http.StatusNotFound, constant.UserNotFound
		}
		return res, http.StatusInternalServerError, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(loginInput.Password))
	if err != nil {
		return res, http.StatusUnauthorized, constant.UsernameOrPasswordInvalid
	}

	var consumerID uuid.UUID
	if len(userData.Consumer) > 0 {
		consumerID = userData.Consumer[0].ID
	}

	jwtClaim := middleware.JwtClaim{
		ID:         userData.ID,
		Email:      userData.Email,
		Role:       constant.Roles(userData.Role.Name),
		ConsumerID: consumerID,
	}

	token, err := getToken(jwtClaim)
	if err != nil {
		return res, http.StatusUnauthorized, err
	}

	localTime, err := utils.FormatTimeWithTimezone(userData.CreatedAt)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = dto.LoginResponse{
		Users: dto.RegisterResponse{
			ID:        userData.ID,
			Email:     userData.Email,
			Role:      userData.Role.Name,
			CreatedAt: localTime,
		},
		Token: token,
	}

	return res, http.StatusOK, nil
}

func getToken(data middleware.JwtClaim) (token string, err error) {
	jwtKey := []byte(config.Env("JWT_SECRET_KEY"))

	expiredDurationStr := config.Env("JWT_EXPIRED_DURATION")
	expiredDuration, err := time.ParseDuration(expiredDurationStr)
	if err != nil {
		return "", constant.InvalidJwtExpiredDurationFormat
	}

	expTime := time.Now().Add(expiredDuration)

	claims := &middleware.JwtClaim{
		ID:         data.ID,
		Email:      data.Email,
		Role:       data.Role,
		ConsumerID: data.ConsumerID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Env("JWT_ISSUER_KEY"),
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenAlgo.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
