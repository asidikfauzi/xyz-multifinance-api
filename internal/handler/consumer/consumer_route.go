package consumer

import (
	"asidikfauzi/xyz-multifinance-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup, consumerController *ConsumersController) {
	consumerGroup := v1.Group("/consumer")
	consumerGroup.Use(middleware.JWTMiddleware())
	{
		consumerGroup.GET("", consumerController.FindAll)
		consumerGroup.GET("/:id", consumerController.FindById)
		consumerGroup.POST("", consumerController.Create)
		consumerGroup.PUT("/:id", consumerController.Update)
	}
}
