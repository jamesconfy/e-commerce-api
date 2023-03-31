package route

import (
	err "e-commerce/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func ErrorRoute(router *gin.Engine) {
	router.NoRoute(err.Error404())
}
