package middleware

import (
	"e-commerce/internal/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type jwtMiddleWare struct {
	tokenSrv service.TokenSrv
}

func (j *jwtMiddleWare) CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid Token")
			return
		}

		auth := strings.Split(authHeader, " ")

		token, err := j.tokenSrv.ValidateToken(auth[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Sprintf("invalid Token: %v", err))
			return
		}

		c.Set("userId", token.Id)
		c.Next()
	}
}

func NewJWTMiddleWare(tokenSrv service.TokenSrv) *jwtMiddleWare {
	return &jwtMiddleWare{tokenSrv: tokenSrv}
}
