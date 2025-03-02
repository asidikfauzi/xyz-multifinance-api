package payment

import (
	"asidikfauzi/xyz-multifinance-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup, paymentController *PaymentsController) {
	paymentGroup := v1.Group("/payment")
	paymentGroup.Use(middleware.JWTMiddleware())
	paymentGroup.Use(middleware.LoggingMiddleware())
	{
		paymentGroup.POST("", paymentController.Create)
	}
}
