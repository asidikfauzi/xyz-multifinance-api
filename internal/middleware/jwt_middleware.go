package middleware

import (
	"asidikfauzi/xyz-multifinance-api/internal/config"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/constant"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type JwtClaim struct {
	ID         uuid.UUID
	Email      string
	Role       constant.Roles
	ConsumerID uuid.UUID
	jwt.RegisteredClaims
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, constant.InvalidHeaderFormat.Error(), nil)
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := &JwtClaim{}

		jwtKey := []byte(config.Env("JWT_SECRET_KEY"))

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrTokenMalformed) {
				response.Error(c, http.StatusUnauthorized, constant.MalformedToken.Error(), nil)
			} else if errors.Is(err, jwt.ErrTokenExpired) {
				response.Error(c, http.StatusUnauthorized, constant.TokenHasExpired.Error(), nil)
			} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
				response.Error(c, http.StatusUnauthorized, constant.TokenIsNotValid.Error(), nil)
			} else {
				response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
			}
			c.Abort()
			return
		}

		if !token.Valid {
			response.Error(c, http.StatusUnauthorized, constant.TokenInvalid.Error(), nil)
			c.Abort()
			return
		}

		fmt.Println(claims)

		c.Set("user_id", claims.ID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Set("consumer_id", claims.ConsumerID)

		c.Next()
	}
}
