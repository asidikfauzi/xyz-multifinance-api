package middleware

import (
	"asidikfauzi/xyz-multifinance-api/internal/config"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")
		expectedApiKey := config.Env("API_KEY")

		if apiKey == "" {
			response.Error(c, http.StatusUnauthorized, constant.APIKeyIsMissing.Error(), nil)
			c.Abort()
			return
		}

		if apiKey != expectedApiKey {
			response.Error(c, http.StatusUnauthorized, constant.InvalidApiKey.Error(), nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
