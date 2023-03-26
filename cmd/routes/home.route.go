package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/internal/service/homeService"

	"github.com/gin-gonic/gin"
)

func HomeRoute(router *gin.RouterGroup, homeSrv homeService.HomeService) {
	handler := handler.NewHomeHandler(homeSrv)
	home := router.Group("/")
	{
		home.GET("", handler.Home)
	}
}
