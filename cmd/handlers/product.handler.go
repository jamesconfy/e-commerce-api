package handler

import (
	"e-commerce/internal/models"
	"e-commerce/internal/models/responseModels"
	"e-commerce/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	AddProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	GetProduct(c *gin.Context)
	EditProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	AddRating(c *gin.Context)
}

type productHanlder struct {
	productSrv service.ProductService
}

func (p *productHanlder) AddProduct(c *gin.Context) {
	var req productModels.AddProductReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
		return
	}

	if err := p.productSrv.Validate(req); err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	req.UserId = c.GetString("userId")
	if req.UserId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responseModels.BuildErrorResponse(http.StatusUnauthorized, "You are not authorized to do that", nil, nil))
		return
	}

	product, err := p.productSrv.AddProduct(&req)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Product added successfully", product, nil))
}

func (p *productHanlder) GetProducts(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}

	pagei, errp := strconv.Atoi(page)
	if errp != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when converting string page to integer", errp, nil))
		return
	}

	products, err := p.productSrv.GetProducts(pagei)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Products fetched successfully", products, nil))
}

func (p *productHanlder) GetProduct(c *gin.Context) {
	productId := c.Param("product_id")

	product, err := p.productSrv.GetProduct(productId)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Product fetched successfully", product, nil))
}

func (p *productHanlder) EditProduct(c *gin.Context) {
	var req models.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
		return
	}

	if err := p.productSrv.Validate(req); err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	req.Id = c.Param("product_id")
	product, err := p.productSrv.GetProduct(req.Id)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	if product.UserId != c.GetString("userId") {
		c.AbortWithStatusJSON(http.StatusForbidden, responseModels.BuildErrorResponse(http.StatusForbidden, "You are not allowed to edit this resource", nil, nil))
		return
	}

	updatedProduct, err := p.productSrv.EditProduct(&req, product)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Product updated successfully", updatedProduct, nil))
}

func (p *productHanlder) DeleteProduct(c *gin.Context) {
	productId := c.Param("product_id")

	userId := c.GetString("userId")
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responseModels.BuildErrorResponse(http.StatusUnauthorized, "You need to be logged in to access this resource", nil, nil))
		return
	}

	product, err := p.productSrv.GetProduct(productId)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}
	if product.UserId != userId {
		c.AbortWithStatusJSON(http.StatusForbidden, responseModels.BuildErrorResponse(http.StatusForbidden, "You are not authorized to delete this resource", nil, nil))
		return
	}

	err = p.productSrv.DeleteProduct(productId)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Product deleted successfully", product, nil))
}

func (p *productHanlder) AddRating(c *gin.Context) {
	var req models.Rating

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", err, nil))
		return
	}

	if err := p.productSrv.Validate(req); err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	req.ProductId = c.Param("product_id")

	_, err := p.productSrv.GetProduct(req.ProductId)
	if err != nil {
		c.AbortWithStatusJSON(err.ResponseCode, err)
		return
	}

	req.UserId = c.GetString("userId")
	if req.UserId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responseModels.BuildErrorResponse(http.StatusUnauthorized, "You are not authorized to do that", nil, nil))
		return
	}

	err = p.productSrv.VerifyUserRatings(req.UserId, req.ProductId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, responseModels.BuildErrorResponse(http.StatusConflict, "You have already rated this product before, try another product", err, nil))
		return
	}

	rating, err := p.productSrv.AddRating(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when adding rating to product", err, nil))
		return
	}

	c.JSON(http.StatusOK, responseModels.BuildSuccessResponse(http.StatusOK, "Product rated successfully", rating, nil))
}

func NewProductHandler(productSrv service.ProductService) ProductHandler {
	return &productHanlder{productSrv: productSrv}
}
