package routes

import (
	err "benny-foodie/cmd/handlers/errorHandler"

	"github.com/gin-gonic/gin"
)

func ErrorRoute(router *gin.Engine) {
	router.NoRoute(err.Error404())
}
