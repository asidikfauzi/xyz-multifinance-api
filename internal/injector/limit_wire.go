//go:build wireinject
// +build wireinject

package injector

import (
	"asidikfauzi/xyz-multifinance-api/internal/database"
	"asidikfauzi/xyz-multifinance-api/internal/handler/limit"
	"asidikfauzi/xyz-multifinance-api/internal/repository/mysql/consumer"
	limitMySQL "asidikfauzi/xyz-multifinance-api/internal/repository/mysql/limit"
	"github.com/google/wire"
)

func InitializedLimitModule() *limit.LimitsController {
	wire.Build(
		database.InitDatabase,
		limitMySQL.NewLimitsMySQL,
		consumer.NewConsumersMySQL,
		limit.NewLimitsService,
		limit.NewLimitsController,
	)

	return nil
}
