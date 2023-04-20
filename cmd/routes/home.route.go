package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/internal/service"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func HomeRoute(router *gin.RouterGroup, homeSrv service.HomeService, store *persistence.InMemoryStore) {
	handler := handler.NewHomeHandler(homeSrv)
	home := router.Group("/")
	{
		home.GET("", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			handler.Home(ctx)
		}))
	}
}
