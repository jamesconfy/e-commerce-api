package producthandler

import (
	"e-commerce/internal/models/productModels"
	"e-commerce/internal/models/responseModels"
	"e-commerce/internal/service/productService"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	AddProduct(c *gin.Context)
	GetProductById(c *gin.Context)
	AddRating(c *gin.Context)
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

func (p *productHanlder) GetProductById(c *gin.Context) {
	productId := c.Param("product_id")
	if productId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, responseModels.BuildErrorResponse(http.StatusNotFound, "No product id was provided", nil, nil))
		return
	}

	product, err := p.productSrv.GetProductById(productId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when getting product", err, nil))
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Product fetched successfully", product, nil))
}

func (p *productHanlder) AddRating(c *gin.Context) {
	var req productModels.AddRatingsReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
		return
	}

	req.ProductId = c.Param("product_id")
	if req.ProductId == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, responseModels.BuildErrorResponse(http.StatusNotFound, "No product id was provided", nil, nil))
		return
	}

	_, err := p.productSrv.GetProductById(req.ProductId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when getting product", err, nil))
		return
	}

	req.UserId = c.GetString("userId")
	if req.UserId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responseModels.BuildErrorResponse(http.StatusUnauthorized, "You are not authorized to do that", nil, nil))
		return
	}

	err = p.productSrv.VerifyUserRatings(req.UserId, req.ProductId)
	fmt.Println(err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responseModels.BuildErrorResponse(http.StatusProxyAuthRequired, "You have already rated this product before, try another product", err, nil))
		return
	}

	rating, err := p.productSrv.AddRating(&req)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when adding rating to product", err, nil))
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Product rated successfully", rating, nil))
}

func NewProductHandler(productSrv productService.ProductService) ProductHandler {
	return &productHanlder{productSrv: productSrv}
}
