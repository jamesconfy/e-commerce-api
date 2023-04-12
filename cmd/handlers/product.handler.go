package handler

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/response"
	"e-commerce/internal/se"
	"e-commerce/internal/service"
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

// Add Product godoc
// @Summary	Add product route
// @Description	Add a product to the database
// @Tags	Product
// @Accept	json
// @Produce	json
// @Param	request	body	forms.Product	true "Product details"
// @Success	200  {object}  response.SuccessMessage{data=models.Product}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/products [post]
// @Security ApiKeyAuth
func (p *productHanlder) Add(c *gin.Context) {
	var req forms.Product

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	userId := c.GetString("userId")

	product, err := p.productSrv.Add(&req, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Product added successfully", product, 1)
}

// Get All Product godoc
// @Summary	Get all product route
// @Description	Provide page number to fetch products
// @Tags	Product
// @Produce	json
// @Param	page	query	string	false "Page"
// @Success	200  {object}  response.SuccessMessage{data=[]models.Product} "asc"
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/products [get]
func (p *productHanlder) GetAll(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}

	pagei, err := strconv.Atoi(page)
	if err != nil {
		response.Error(c, *se.New("Error when converting string to integer", err, se.ErrServer))
		return
	}

	products, errI := p.productSrv.GetAll(pagei)
	if err != nil {
		response.Error(c, *errI)
		return
	}

	response.Success(c, "Products fetched successfully", products, len(products))
}

// Get Product godoc
// @Summary	Get product route
// @Description	Get a product when provided with the id
// @Tags	Product
// @Produce	json
// @Param	productId	path	string	true "Product Id"
// @Success	200  {object}  response.SuccessMessage{data=models.Product}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/products/:productId [get]
func (p *productHanlder) Get(c *gin.Context) {
	productId := c.Param("product_id")

	product, err := p.productSrv.Get(productId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Product fetched successfully", product, 1)
}

// Edit Product godoc
// @Summary	Edit product route
// @Description	Edit a product by providing both a request and the product id
// @Tags	Product
// @Produce	json
// @Param	productId	path	string	true "Product Id"
// @Param	request	body	forms.EditProduct	true "Edit product request"
// @Success	200  {object}  response.SuccessMessage{data=models.Product}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/products/:productId [patch]
// @Security ApiKeyAuth
func (p *productHanlder) Edit(c *gin.Context) {
	var req forms.EditProduct

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	productId := c.Param("product_id")
	userId := c.GetString("userId")

	product, err := p.productSrv.Edit(&req, productId, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Product updated successfully", product, 1)
}

// Delete Product godoc
// @Summary	Delete product route
// @Description	Delete a product by it's id
// @Tags	Product
// @Produce	json
// @Param	productId	path	string	true "Product Id"
// @Success	200  {string}	string	true	"Product deleted successfully"
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/products/:productId [delete]
// @Security ApiKeyAuth
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

// Add Rating godoc
// @Summary	Add rating route
// @Description	Add a rating to a product
// @Tags	Product
// @Produce	json
// @Param	productId	path	string	true "Product Id"
// @Param	request	body	forms.Rating	true "Add rating request"
// @Success	200  {object}  response.SuccessMessage{data=models.Rating}
// @Failure	400  {object}  response.ErrorMessage
// @Failure	404  {object}  response.ErrorMessage
// @Failure	500  {object}  response.ErrorMessage
// @Router	/products/:productId/ratings [post]
// @Security ApiKeyAuth
func (p *productHanlder) AddRating(c *gin.Context) {
	var req forms.Rating

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, *se.Validating(err))
		return
	}

	productId := c.Param("product_id")
	userId := c.GetString("userId")

	rating, err := p.productSrv.AddRating(&req, productId, userId)
	if err != nil {
		response.Error(c, *err)
		return
	}

	response.Success(c, "Product rated successfully", rating)
}

func NewProductHandler(productSrv service.ProductService) ProductHandler {
	return &productHanlder{productSrv: productSrv}
}
