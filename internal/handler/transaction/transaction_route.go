package transaction

import (
	"asidikfauzi/xyz-multifinance-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup, transactionController *TransactionsController) {
	limitGroup := v1.Group("/transaction")
	limitGroup.Use(middleware.JWTMiddleware())
	limitGroup.Use(middleware.LoggingMiddleware())
	{
		limitGroup.POST("", transactionController.Transactions)
	}
}
