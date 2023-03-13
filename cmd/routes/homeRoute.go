package routes

import (
	"e-commerce/cmd/handlers/homeHandler"
	"e-commerce/internal/service/homeService"

	"github.com/gin-gonic/gin"
)

func HomeRoute(router *gin.RouterGroup, homeSrv homeService.HomeService) {
	handler := homeHandler.NewHomeHandler(homeSrv)
	home := router.Group("/")
	{
		home.GET("", handler.Home)
	}
}
