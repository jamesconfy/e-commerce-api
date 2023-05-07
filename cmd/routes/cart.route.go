package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func CartRoute(router *gin.RouterGroup, cartSrv service.CartService, authSrv service.AuthService, store *persistence.InMemoryStore) {
	handler := handler.NewCartHandler(cartSrv)
	auth := middleware.Authentication(authSrv)
	cart := router.Group("/carts")

	cart.Use(auth.CheckJWT())
	{
		cart.GET("", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			handler.Get(ctx)
		}))
		cart.DELETE("", func(ctx *gin.Context) {
			defer func() {
				store.Delete(cache.CreateKey("/api/v1/items"))
			}()
			handler.Clear(ctx)
		})
	}
}
