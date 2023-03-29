package handler

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/response"
	"e-commerce/internal/service"
	"e-commerce/internal/serviceerror"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	Add(c *gin.Context)
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Edit(c *gin.Context)
	Delete(c *gin.Context)
	AddRating(c *gin.Context)
}

type productHanlder struct {
	productSrv service.ProductService
}

func (p *productHanlder) Add(c *gin.Context) {
	var req forms.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *serviceerror.Validating(err))
		return
	}

	userId := c.GetString("userId")

	product, err := p.productSrv.Add(&req, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Product added successfully", product)
}

func (p *productHanlder) GetAll(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}

	pagei, err := strconv.Atoi(page)
	if err != nil {
		response.Error(c, *serviceerror.New("Error when converting string to integer", err, serviceerror.ErrServer))
		return
	}

	products, errI := p.productSrv.GetAll(pagei)
	if err != nil {
		response.Error(c, *errI)
		return
	}

	response.Success(c, "Products fetched successfully", products)
}

func (p *productHanlder) Get(c *gin.Context) {
	productId := c.Param("product_id")

	product, err := p.productSrv.Get(productId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Product fetched successfully", product, nil)
}

func (p *productHanlder) Edit(c *gin.Context) {
	var req forms.EditProduct

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *serviceerror.Validating(err))
		return
	}

	productId := c.Param("product_id")
	userId := c.GetString("userId")

	product, err := p.productSrv.Edit(&req, productId, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Product updated successfully", product, nil)
}

func (p *productHanlder) Delete(c *gin.Context) {
	productId := c.Param("product_id")
	userId := c.GetString("userId")

	err := p.productSrv.Delete(productId, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success202(c, "Product deleted successfully")
}

func (p *productHanlder) AddRating(c *gin.Context) {
	var req forms.Rating

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *serviceerror.Validating(err))
		return
	}

	req.ProductId = c.Param("product_id")
	userId := c.GetString("userId")

	rating, err := p.productSrv.AddRating(&req, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Product rated successfully", rating)
}

func NewProductHandler(productSrv service.ProductService) ProductHandler {
	return &productHanlder{productSrv: productSrv}
}
