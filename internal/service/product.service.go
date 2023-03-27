package service

// import (
// 	"e-commerce/internal/models/errorModels"
// 	"e-commerce/internal/models/productModels"
// 	"e-commerce/internal/models/responseModels"
// 	repo "e-commerce/internal/repository"
// 	"e-commerce/utils"
// 	"net/http"

// 	"github.com/google/uuid"
// )

// type ProductService interface {
// 	Validate(req any) *responseModels.ResponseMessage
// 	AddProduct(req *productModels.AddProductReq) (*productModels.AddProductRes, *responseModels.ResponseMessage)
// 	GetProducts(page int) ([]*productModels.GetProductRes, *responseModels.ResponseMessage)
// 	GetProduct(productId string) (*productModels.GetProductRes, *responseModels.ResponseMessage)
// 	EditProduct(req *productModels.EditProductReq, product *productModels.GetProductRes) (*productModels.EditProductRes, *responseModels.ResponseMessage)
// 	DeleteProduct(productId string) *responseModels.ResponseMessage
// 	AddRating(req *productModels.AddRatingsReq) (*productModels.AddRatingsRes, *responseModels.ResponseMessage)
// 	VerifyUserRatings(userId, productId string) *responseModels.ResponseMessage
// }

// type productSrv struct {
// 	repo         repo.ProductRepo
// 	validatorSrv ValidationSrv
// 	loggerSrv    LogSrv
// 	timeSrv      TimeService
// 	message      utils.Messages
// }

// func (p *productSrv) Validate(req any) *responseModels.ResponseMessage {
// 	if err := p.validatorSrv.Validate(req); err != nil {
// 		e := errorModels.NewValidatingError(err)
// 		return responseModels.BuildErrorResponse(http.StatusBadRequest, "Bad input data", e, nil)
// 	}

// 	return nil
// }

// func (p *productSrv) AddProduct(req *productModels.AddProductReq) (*productModels.AddProductRes, *responseModels.ResponseMessage) {
// 	req.DateCreated = p.timeSrv.CurrentTime()
// 	req.DateUpdated = p.timeSrv.CurrentTime()
// 	req.ProductId = uuid.New().String()

// 	err := p.repo.Add(req)
// 	if err != nil {
// 		// p.loggerSrv.Fatal(p.message.AddProductRepoError(req, err))
// 		return nil, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when adding product to database", err, nil)
// 	}

// 	data := &productModels.AddProductRes{
// 		ProductId:   req.ProductId,
// 		Name:        req.Name,
// 		Description: req.Description,
// 		Price:       req.Price,
// 		Image:       req.Image,
// 	}

// 	// p.loggerSrv.Info(p.message.AddProductSuccess(req))
// 	return data, nil
// }

// func (p *productSrv) GetProducts(page int) ([]*productModels.GetProductRes, *responseModels.ResponseMessage) {
// 	products, err := p.repo.GetAll(page)
// 	if err != nil {
// 		// p.loggerSrv.Fatal(p.message.InternalServerError(err))
// 		return nil, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when getting products", err, nil)
// 	}

// 	// p.loggerSrv.Info(p.message.GetProductsSuccess())
// 	return products, nil
// }

// func (p *productSrv) GetProduct(productId string) (*productModels.GetProductRes, *responseModels.ResponseMessage) {
// 	product, err := p.repo.GetId(productId)
// 	if product != nil && err != nil {
// 		// p.loggerSrv.Fatal(p.message.GetProductNotFound(productId, err))
// 		return product, responseModels.BuildErrorResponse(http.StatusNotFound, "Product not found", err, nil)
// 	}

// 	if product == nil && err != nil {
// 		// p.loggerSrv.Fatal(p.message.InternalServerError(err))
// 		return nil, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Internal server error", err, nil)
// 	}

// 	// p.loggerSrv.Info(p.message.GetProductSuccess(product))
// 	return product, nil
// }

// func (p *productSrv) EditProduct(req *productModels.EditProductReq, product *productModels.GetProductRes) (*productModels.EditProductRes, *responseModels.ResponseMessage) {
// 	editProduct := p.updateProduct(*req, product)

// 	err := p.repo.Edit(editProduct)
// 	if err != nil {
// 		// p.loggerSrv.Fatal(p.message.InternalServerError(err))
// 		return nil, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Internal server error", err, nil)
// 	}

// 	// p.loggerSrv.Info(p.message.EditProductSuccess(editProduct))
// 	return &productModels.EditProductRes{
// 		ProductId:   editProduct.ProductId,
// 		Name:        editProduct.Name,
// 		Description: editProduct.Description,
// 		Price:       editProduct.Price,
// 		DateUpdated: editProduct.DateUpdated,
// 		Image:       editProduct.Image,
// 	}, nil
// }

// func (p *productSrv) DeleteProduct(productId string) *responseModels.ResponseMessage {
// 	err := p.repo.Delete(productId)
// 	if err != nil {
// 		p.loggerSrv.Fatal(p.message.DeleteProductRepoError(productId, err))
// 		return responseModels.BuildErrorResponse(http.StatusInternalServerError, "Internal server error", err, nil)
// 	}

// 	p.loggerSrv.Info(p.message.DeleteProductSuccess(productId))
// 	return nil
// }

// func (p *productSrv) AddRating(req *productModels.AddRatingsReq) (*productModels.AddRatingsRes, *responseModels.ResponseMessage) {
// 	req.RatingId = uuid.New().String()
// 	req.DateCreated = p.timeSrv.CurrentTime()
// 	req.DateUpdated = p.timeSrv.CurrentTime()

// 	err := p.repo.AddRating(req)
// 	if err != nil {
// 		p.loggerSrv.Fatal(p.message.AddRatingRepoError(req))
// 		return nil, responseModels.BuildErrorResponse(http.StatusInternalServerError, "Error when rating product", err, nil)
// 	}

// 	data := &productModels.AddRatingsRes{
// 		RatingId:    req.RatingId,
// 		Rating:      req.Rating,
// 		ProductId:   req.ProductId,
// 		UserId:      req.UserId,
// 		DateCreated: req.DateCreated,
// 		DateUpdated: req.DateUpdated,
// 	}

// 	p.loggerSrv.Info(p.message.AddRatingSuccess(data))
// 	return data, nil
// }

// func (p *productSrv) VerifyUserRatings(userId, productId string) *responseModels.ResponseMessage {
// 	err := p.repo.VerifyRating(userId, productId)
// 	if err != nil {
// 		// p.loggerSrv.Error(p.message.VerifyUserRatingsRepoError(userId, productId))
// 		return responseModels.BuildErrorResponse(http.StatusConflict, "You cannot re-rate this product", err, nil)
// 	}

// 	// p.loggerSrv.Info(p.message.VerifyUserRatingsSucess(userId, productId))
// 	return nil
// }

// func NewProductService(productRepo repo.ProductRepo, validatorSrv ValidationSrv, loggerSrv LogSrv, timeSrv TimeService, message utils.Messages) ProductService {
// 	return &productSrv{repo: productRepo, validatorSrv: validatorSrv, loggerSrv: loggerSrv, timeSrv: timeSrv, message: message}
// }

// // Auxillary Function
// func (p *productSrv) updateProduct(req productModels.EditProductReq, product *productModels.GetProductRes) *productModels.EditProductReq {
// 	if req.Name != "" && req.Name != product.Name {
// 		product.Name = req.Name
// 	}

// 	if req.Description != "" && req.Description != product.Description {
// 		product.Description = req.Description
// 	}

// 	if req.Price != 0.0 && req.Price != product.Price {
// 		product.Price = req.Price
// 	}

// 	if req.Image != "" && req.Image != product.Image {
// 		product.Image = req.Image
// 	}

// 	product.DateUpdated = p.timeSrv.CurrentTime()

// 	return &productModels.EditProductReq{
// 		Name:        product.Name,
// 		Description: product.Description,
// 		ProductId:   product.ProductId,
// 		UserId:      product.UserId,
// 		DateUpdated: product.DateUpdated,
// 		Price:       product.Price,
// 		Image:       product.Image,
// 	}
// }
