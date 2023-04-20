package route

import (
	handler "e-commerce/cmd/handlers"
	"e-commerce/cmd/middleware"
	"e-commerce/internal/service"
	"fmt"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(v1 *gin.RouterGroup, productSrv service.ProductService, authSrv service.AuthService, store *persistence.InMemoryStore) {
	handler := handler.NewProductHandler(productSrv)
	auth := middleware.Authentication(authSrv)

	product := v1.Group("/products")
	product.Use(auth.CheckJWT())
	{
		product.POST("", func(ctx *gin.Context) {
			defer func() {
				store.Delete(cache.CreateKey("/api/v1/products"))
			}()

			handler.Add(ctx)
		})
		product.PATCH("/:productId", func(ctx *gin.Context) {
			defer func() {
				store.Delete(cache.CreateKey(ctx.Request.RequestURI))
				store.Delete(cache.CreateKey("/api/v1/products"))
			}()

			handler.Edit(ctx)
		})
		product.DELETE("/:productId", func(ctx *gin.Context) {
			defer func() {
				store.Delete(cache.CreateKey("/api/v1/products"))
				key := fmt.Sprintf("/api/v1/products/%v", ctx.Param("productId"))
				store.Delete(cache.CreateKey(key))
			}()

			handler.Delete(ctx)
		})
		product.POST("/:productId/ratings", func(ctx *gin.Context) {
			defer func() {
				store.Delete(cache.CreateKey("/api/v1/products"))
				key := fmt.Sprintf("/api/v1/products/%v", ctx.Param("productId"))
				store.Delete(cache.CreateKey(key))
			}()

			handler.AddRating(ctx)
		})
	}

	product1 := v1.Group("/products")
	{
		product1.GET("/:productId", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			handler.Get(ctx)
		}))
		product1.GET("", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			handler.GetAll(ctx)
		}))
	}
}
