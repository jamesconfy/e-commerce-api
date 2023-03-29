package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	se "e-commerce/internal/serviceerror"
	"e-commerce/utils"

	"github.com/google/uuid"
)

type ProductService interface {
	Validate(req any) error
	Add(req *forms.Product, userId string) (*models.Product, *se.ServiceError)
	GetAll(page int) ([]*models.ProductRating, *se.ServiceError)
	Get(productId string) (*models.ProductRating, *se.ServiceError)
	Edit(req *forms.EditProduct, productId, userId string) (*models.Product, *se.ServiceError)
	Delete(productId, userId string) *se.ServiceError
	AddRating(req *forms.Rating, userId string) (*models.Rating, *se.ServiceError)
	// VerifyUserRatings(userId, productId string) *responseModels.ResponseMessage
}

type productSrv struct {
	repo         repo.ProductRepo
	validatorSrv ValidationSrv
	loggerSrv    LogSrv
	timeSrv      TimeService
	message      utils.Messages
}

func (p *productSrv) Validate(req any) error {
	err := p.validatorSrv.Validate(req)
	if err != nil {
		return err
	}

	return nil
}

func (p *productSrv) Add(req *forms.Product, userId string) (*models.Product, *se.ServiceError) {
	if err := p.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	var product models.Product

	product.Id = uuid.New().String()
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Image = req.Image
	product.UserId = userId
	product.DateCreated = p.timeSrv.CurrentTime()
	product.DateUpdated = p.timeSrv.CurrentTime()

	result, err := p.repo.Add(&product)
	if err != nil {
		// p.loggerSrv.Fatal(p.message.AddProductRepoError(req, err))
		return nil, se.Internal(err)
	}

	// p.loggerSrv.Info(p.message.AddProductSuccess(req))
	return result, nil
}

func (p *productSrv) GetAll(page int) ([]*models.ProductRating, *se.ServiceError) {
	products, err := p.repo.GetAll(page)
	if err != nil && products == nil {
		// p.loggerSrv.Fatal(p.message.InternalServerError(err))
		return nil, se.Internal(err)
	}

	// p.loggerSrv.Info(p.message.GetProductsSuccess())
	return products, nil
}

func (p *productSrv) Get(productId string) (*models.ProductRating, *se.ServiceError) {
	product, err := p.repo.GetId(productId)
	// p.loggerSrv.Fatal(p.message.GetProductNotFound(productId, err))

	if err != nil {
		// p.loggerSrv.Fatal(p.message.InternalServerError(err))
		return nil, se.NotFoundOrInternal(err)
	}

	// p.loggerSrv.Info(p.message.GetProductSuccess(product))
	return product, nil
}

func (p *productSrv) Edit(req *forms.EditProduct, productId, userId string) (*models.Product, *se.ServiceError) {
	if err := p.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	product, err := p.repo.GetId(productId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err)
	}

	if product.Product.UserId != userId {
		return nil, se.Forbidden("You are not able to modify that resource")
	}

	editProduct := p.updateProduct(req, product)

	returnProduct, err := p.repo.Edit(editProduct)
	if err != nil {
		// p.loggerSrv.Fatal(p.message.InternalServerError(err))
		return nil, se.Internal(err)
	}

	// p.loggerSrv.Info(p.message.EditProductSuccess(editProduct))
	return returnProduct, nil
}

func (p *productSrv) Delete(productId, userId string) *se.ServiceError {
	product, err := p.repo.GetId(productId)
	if err != nil {
		return se.NotFoundOrInternal(err)
	}

	if product.Product.UserId != userId {
		return se.Forbidden("You are not able to modify that resource")
	}

	err = p.repo.Delete(productId)
	if err != nil {
		// p.loggerSrv.Fatal(p.message.DeleteProductRepoError(productId, err))
		return se.Internal(err)
	}

	// p.loggerSrv.Info(p.message.DeleteProductSuccess(productId))
	return nil
}

func (p *productSrv) AddRating(req *forms.Rating, userId string) (*models.Rating, *se.ServiceError) {
	if err := p.Validate(req); err != nil {
		return nil, se.Internal(err)
	}

	product, err := p.repo.GetId(req.ProductId)
	if err != nil {
		return nil, se.NotFoundOrInternal(err)
	}

	if product.Product.UserId == userId {
		return nil, se.Forbidden("You cannot rate your own product")
	}

	var rating models.Rating

	rating.Value = req.Value
	rating.ProductId = req.ProductId
	rating.UserId = userId
	rating.DateCreated = p.timeSrv.CurrentTime()
	rating.DateUpdated = p.timeSrv.CurrentTime()

	result, err := p.repo.AddRating(&rating)
	if err != nil {
		// p.loggerSrv.Fatal(p.message.AddRatingRepoError(req))
		return nil, se.Internal(err)
	}

	// p.loggerSrv.Info(p.message.AddRatingSuccess(data))
	return result, nil
}

// func (p *productSrv) VerifyUserRatings(userId, productId string) *responseModels.ResponseMessage {
// 	err := p.repo.VerifyRating(userId, productId)
// 	if err != nil {
// 		// p.loggerSrv.Error(p.message.VerifyUserRatingsRepoError(userId, productId))
// 		return responseModels.BuildErrorResponse(http.StatusConflict, "You cannot re-rate this product", err, nil)
// 	}

// 	// p.loggerSrv.Info(p.message.VerifyUserRatingsSucess(userId, productId))
// 	return nil
// }

func NewProductService(productRepo repo.ProductRepo, validatorSrv ValidationSrv, loggerSrv LogSrv, timeSrv TimeService, message utils.Messages) ProductService {
	return &productSrv{repo: productRepo, validatorSrv: validatorSrv, loggerSrv: loggerSrv, timeSrv: timeSrv, message: message}
}

// Auxillary Function
func (p *productSrv) updateProduct(req *forms.EditProduct, product *models.ProductRating) *models.Product {
	if req.Name != "" && req.Name != product.Product.Name {
		product.Product.Name = req.Name
	}

	if req.Description != "" && req.Description != product.Product.Description {
		product.Product.Description = req.Description
	}

	if req.Price != 0.0 && req.Price != product.Product.Price {
		product.Product.Price = req.Price
	}

	if req.Image != "" && req.Image != product.Product.Image {
		product.Product.Image = req.Image
	}

	product.Product.DateUpdated = p.timeSrv.CurrentTime()

	return &models.Product{
		Id:          product.Product.Id,
		UserId:      product.Product.UserId,
		Name:        product.Product.Name,
		Description: product.Product.Description,
		Price:       product.Product.Price,
		Image:       product.Product.Image,
		DateUpdated: product.Product.DateUpdated,
	}
}
