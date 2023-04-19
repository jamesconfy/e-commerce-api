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
		product.PATCH("/:product_id", handler.Edit)
		product.DELETE("/:product_id", handler.Delete)
		product.POST("/:product_id/ratings", handler.AddRating)
	}

	product1 := v1.Group("/products")
	{
		product1.GET("/:product_id", handler.Get)
		product1.GET("", cache.CachePage(store, time.Duration(time.Hour*1), func(ctx *gin.Context) {
			fmt.Println(ctx.Request.RequestURI)
			handler.GetAll(ctx)
		}))
	}
}
