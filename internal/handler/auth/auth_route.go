package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(v1 *gin.RouterGroup, authController *AuthsController) {
	authGroup := v1.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/register", authController.Register)
	}
}
