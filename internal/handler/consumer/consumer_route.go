package consumer

import (
	"asidikfauzi/xyz-multifinance-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup, consumerController *ConsumersController) {
	consumerGroup := v1.Group("/consumer")
	consumerGroup.Use(middleware.JWTMiddleware())
	consumerGroup.Use(middleware.LoggingMiddleware())
	{
		consumerGroup.GET("", consumerController.FindAll)
		consumerGroup.GET("/:id", consumerController.FindById)
		consumerGroup.PUT("/:id", consumerController.Update)
	}
}
