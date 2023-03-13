package middleware

import (
	tokenservice "e-commerce/internal/service/tokenService"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type jwtMiddleWare struct {
	tokenSrv tokenservice.TokenSrv
}

func (j *jwtMiddleWare) ValidateJWT() gin.HandlerFunc {
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

func NewJWTMiddleWare(tokenSrv tokenservice.TokenSrv) *jwtMiddleWare {
	return &jwtMiddleWare{tokenSrv: tokenSrv}
}
