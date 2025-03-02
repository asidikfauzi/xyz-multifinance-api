//go:build wireinject
// +build wireinject

package injector

import (
	"asidikfauzi/xyz-multifinance-api/internal/database"
	"asidikfauzi/xyz-multifinance-api/internal/handler/consumer"
	consumerMySQL "asidikfauzi/xyz-multifinance-api/internal/repository/mysql/consumer"
	"github.com/google/wire"
)

func InitializedConsumerModule() *consumer.ConsumersController {
	wire.Build(
		database.InitDatabase,
		consumerMySQL.NewConsumersMySQL,
		consumer.NewConsumersService,
		consumer.NewConsumersController,
	)

	return nil
}
