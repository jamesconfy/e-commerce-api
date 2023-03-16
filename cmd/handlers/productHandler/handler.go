package producthandler

import (
	"e-commerce/internal/models/productModels"
	"e-commerce/internal/models/responseModels"
	"e-commerce/internal/service/productService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	AddProduct(c *gin.Context)
}

type productHanlder struct {
	productSrv productService.ProductService
}

func (p *productHanlder) AddProduct(c *gin.Context) {
	var req productModels.AddProductReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
		return
	}

	req.UserId = c.GetString("userId")
	if req.UserId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responseModels.BuildErrorResponse(http.StatusUnauthorized, "You are not authorized to do that", nil, nil))
		return
	}

	product, errP := p.productSrv.AddProduct(&req)
	if errP != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error creating user", errP, nil))
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Product added successfully", product, nil))
}

func NewProductHandler(productSrv productService.ProductService) ProductHandler {
	return &productHanlder{productSrv: productSrv}
}
