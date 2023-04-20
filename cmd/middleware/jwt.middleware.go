package middleware

import (
	"e-commerce/internal/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWT interface {
	CheckJWT() gin.HandlerFunc
}

type jwtMiddleWare struct {
	tokenSrv service.AuthService
}

func (j *jwtMiddleWare) CheckJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := GetAuthorizationHeader(c)
		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Authorization Token: Token cannot be empty"})
			return
		}

		token, err := j.tokenSrv.Validate(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Sprintf("Invalid Authorization Token: %v", err))
			return
		}

		c.Set("userId", token.Id)
		c.Next()
	}
}

func GetAuthorizationHeader(c *gin.Context) string {
	if isBrowser(c.Request.UserAgent()) {
		authtoken, _ := c.Cookie("Authorization")
		return authtoken
	}

	authHeader := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
	return authHeader
}

func isBrowser(userAgent string) bool {
	switch {
	case strings.Contains(userAgent, "Mozilla"), strings.Contains(userAgent, "Chrome"), strings.Contains(userAgent, "Postman"), strings.Contains(userAgent, "Edge"), strings.Contains(userAgent, "Trident"):
		return true
	default:
		return false
	}
}

func Authentication(authSrv service.AuthService) JWT {
	return &jwtMiddleWare{tokenSrv: authSrv}
}
