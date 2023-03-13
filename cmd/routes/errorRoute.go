package routes

import (
	err "e-commerce/cmd/handlers/errorHandler"

	"github.com/gin-gonic/gin"
)

func ErrorRoute(router *gin.Engine) {
	router.NoRoute(err.Error404())
}
