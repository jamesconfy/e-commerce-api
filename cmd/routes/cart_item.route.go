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

func CartItemRoute(router *gin.RouterGroup, cartItemSrv service.CartItemService, authSrv service.AuthService, store *persistence.InMemoryStore) {
	handler := handler.NewCartItemHandler(cartItemSrv)
	auth := middleware.Authentication(authSrv)

	item := router.Group("/items")
	item.Use(auth.CheckJWT())
	{
		item.POST("", func(ctx *gin.Context) {
			store.Delete(cache.CreateKey("/api/v1/items"))
			handler.Add(ctx)
		})
		item.GET("", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			handler.GetItems(ctx)
		}))
		item.GET("/:productId", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			handler.Get(ctx)
		}))
		item.DELETE("/:productId", func(ctx *gin.Context) {
			defer func() {
				store.Delete(cache.CreateKey(ctx.Request.RequestURI))
				store.Delete(cache.CreateKey("/api/v1/items"))
			}()
			handler.Delete(ctx)
		})
	}
}
