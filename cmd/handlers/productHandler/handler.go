package producthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	Product(c *gin.Context)
}

type productHanlder struct{}

func (p productHanlder) Product(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello from product handler")
}

func NewProductHandler() ProductHandler {
	return &productHanlder{}
}
