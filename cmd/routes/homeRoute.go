package routes

import (
	"benny-foodie/cmd/handlers/homeHandler"
	"benny-foodie/internal/service/homeService"

	"github.com/gin-gonic/gin"
)

func HomeRoute(router *gin.RouterGroup, homeSrv homeService.HomeService) {
	handler := homeHandler.NewHomeHandler(homeSrv)
	home := router.Group("/")
	{
		home.GET("", handler.Home)
	}
}
