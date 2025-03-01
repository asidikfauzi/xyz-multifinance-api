//go:build wireinject
// +build wireinject

package injector

import (
	"asidikfauzi/xyz-multifinance-api/internal/database"
	"asidikfauzi/xyz-multifinance-api/internal/handler/auth"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/role"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/user"
	"github.com/google/wire"
)

func InitializedAuthModule() *auth.AuthsController {
	wire.Build(
		database.InitDatabase,
		user.NewUserMySQL,
		role.NewRolesMySQL,
		auth.NewAuthService,
		auth.NewAuthController,
	)

	return nil
}
