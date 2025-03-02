// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"asidikfauzi/xyz-multifinance-api/internal/database"
	"asidikfauzi/xyz-multifinance-api/internal/handler/auth"
	"asidikfauzi/xyz-multifinance-api/internal/handler/consumer"
	"asidikfauzi/xyz-multifinance-api/internal/handler/limit"
	consumer2 "asidikfauzi/xyz-multifinance-api/internal/repository/mysql/consumer"
	limit2 "asidikfauzi/xyz-multifinance-api/internal/repository/mysql/limit"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/role"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/user"
)

// Injectors from auth_wire.go:

func InitializedAuthModule() *auth.AuthsController {
	db := database.InitDatabase()
	usersMySQL := user.NewUserMySQL(db)
	rolesMySQL := role.NewRolesMySQL(db)
	authsService := auth.NewAuthService(usersMySQL, rolesMySQL)
	authsController := auth.NewAuthController(authsService)
	return authsController
}

// Injectors from consumer_wire.go:

func InitializedConsumerModule() *consumer.ConsumersController {
	db := database.InitDatabase()
	consumersMySQL := consumer2.NewConsumersMySQL(db)
	consumersService := consumer.NewConsumersService(consumersMySQL)
	consumersController := consumer.NewConsumersController(consumersService)
	return consumersController
}

// Injectors from limit_wire.go:

func InitializedLimitModule() *limit.LimitsController {
	db := database.InitDatabase()
	limitsMySQL := limit2.NewLimitsMySQL(db)
	consumersMySQL := consumer2.NewConsumersMySQL(db)
	limitsService := limit.NewLimitsService(limitsMySQL, consumersMySQL)
	limitsController := limit.NewLimitsController(limitsService)
	return limitsController
}
