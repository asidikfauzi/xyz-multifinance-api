package limit

import (
	"asidikfauzi/xyz-multifinance-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(v1 *gin.RouterGroup, limitController *LimitsController) {
	limitGroup := v1.Group("/limit")
	limitGroup.Use(middleware.JWTMiddleware())
	{
		limitGroup.PATCH("/approval/:id", limitController.ApprovalConsumer)
	}
}
