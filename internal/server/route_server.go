package server

import (
	"asidikfauzi/xyz-multifinance-api/internal/handler/auth"
	"asidikfauzi/xyz-multifinance-api/internal/handler/consumer"
	"asidikfauzi/xyz-multifinance-api/internal/handler/limit"
	"asidikfauzi/xyz-multifinance-api/internal/handler/transaction"
	"asidikfauzi/xyz-multifinance-api/internal/injector"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func InitializedServer() *Server {
	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		fmt.Println("pong")
		c.String(200, "pong")
	})

	authModule := injector.InitializedAuthModule()
	auth.RegisterRoutes(v1, authModule)

	consumerModule := injector.InitializedConsumerModule()
	consumer.RegisterRoutes(v1, consumerModule)

	limitModule := injector.InitializedLimitModule()
	limit.RegisterRoutes(v1, limitModule)

	transactionModule := injector.InitializedTransactionModule()
	transaction.RegisterRoutes(v1, transactionModule)

	return &Server{Engine: r}
}
