package service

import (
	"e-commerce/internal/forms"
	"e-commerce/internal/logger"
	"e-commerce/internal/models"
	repo "e-commerce/internal/repository"
	se "e-commerce/internal/se"
	"time"

	"github.com/google/uuid"
)

type ProductService interface {
	Validate(req any) error
	Add(req *forms.Product, userId string) (*models.Product, *se.ServiceError)
	GetAll(page int) ([]*models.ProductRating, *se.ServiceError)
	Get(productId string) (*models.ProductRating, *se.ServiceError)
	Edit(req *forms.EditProduct, productId, userId string) (*models.Product, *se.ServiceError)
	Delete(productId, userId string) *se.ServiceError
	AddRating(req *forms.Rating, productId, userId string) (*models.Rating, *se.ServiceError)
}

type productSrv struct {
	productRepo  repo.ProductRepo
	validatorSrv ValidationSrv
	loggerSrv    LogSrv
	message      logger.Messages
}

func (p *productSrv) Validate(req any) error {
	err := p.validatorSrv.Validate(req)
	if err != nil {
		p.loggerSrv.Error(p.message.ValidationError(req, err))
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

	result, err := p.productRepo.Add(&product)
	if err != nil {
		p.loggerSrv.Fatal(p.message.AddRepoError(&product, err))
		return nil, se.Internal(err)
	}

	p.loggerSrv.Info(p.message.AddSuccess(result))
	return result, nil
}

func (p *productSrv) GetAll(page int) ([]*models.ProductRating, *se.ServiceError) {
	products, err := p.productRepo.GetAll(page)
	if err != nil && products == nil {
		p.loggerSrv.Fatal(p.message.GetAllRepoError(err))
		return nil, se.Internal(err)
	}

	p.loggerSrv.Info(p.message.GetAllSuccess(products))
	return products, nil
}

func (p *productSrv) Get(productId string) (*models.ProductRating, *se.ServiceError) {
	product, err := p.productRepo.Get(productId)

	if err != nil {
		p.loggerSrv.Fatal(p.message.GetProductRepoError(productId, err))
		return nil, se.NotFoundOrInternal(err)
	}

	p.loggerSrv.Info(p.message.GetProductSuccess(product.Product))
	return product, nil
}

func (p *productSrv) Edit(req *forms.EditProduct, productId, userId string) (*models.Product, *se.ServiceError) {
	if err := p.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	product, err := p.productRepo.Get(productId)
	if err != nil {
		p.loggerSrv.Fatal(p.message.GetProductRepoError(productId, err))
		return nil, se.NotFoundOrInternal(err)
	}

	if product.Product.UserId != userId {
		p.loggerSrv.Fatal(p.message.EditCompareUser(product.Product.UserId, userId))
		return nil, se.Forbidden("You are not able to modify that resource")
	}

	editProduct := p.updateProduct(req, product)

	returnProduct, err := p.productRepo.Edit(editProduct)
	if err != nil {
		p.loggerSrv.Fatal(p.message.EditProductRepoError(editProduct, err))
		return nil, se.Internal(err)
	}

	p.loggerSrv.Info(p.message.EditProductSuccess(editProduct))
	return returnProduct, nil
}

func (p *productSrv) Delete(productId, userId string) *se.ServiceError {
	product, err := p.productRepo.Get(productId)
	if err != nil {
		p.loggerSrv.Error(p.message.GetProductRepoError(productId, err))
		return se.NotFoundOrInternal(err)
	}

	if product.Product.UserId != userId {
		p.loggerSrv.Fatal(p.message.EditCompareUser(product.Product.UserId, userId))
		return se.Forbidden("You are not able to modify that resource")
	}

	err = p.productRepo.Delete(productId)
	if err != nil {
		p.loggerSrv.Fatal(p.message.DeleteProductRepoError(productId, err))
		return se.Internal(err)
	}

	p.loggerSrv.Info(p.message.DeleteProductSuccess(productId))
	return nil
}

func (p *productSrv) AddRating(req *forms.Rating, productId, userId string) (*models.Rating, *se.ServiceError) {
	if err := p.Validate(req); err != nil {
		return nil, se.Validating(err)
	}

	product, err := p.productRepo.Get(productId)
	if err != nil {
		p.loggerSrv.Error(p.message.GetProductRepoError(productId, err))
		return nil, se.NotFoundOrInternal(err)
	}

	if product.Product.UserId == userId {
		p.loggerSrv.Fatal(p.message.AddRatingCompareUser(req, product.Product.UserId, userId))
		return nil, se.Forbidden("You cannot rate your own product")
	}

	var rating models.Rating

	rating.Value = req.Value
	rating.ProductId = productId
	rating.UserId = userId

	result, err := p.productRepo.AddRating(&rating)
	if err != nil {
		p.loggerSrv.Fatal(p.message.AddRatingRepoError(&rating, err))
		return nil, se.Internal(err)
	}

	p.loggerSrv.Info(p.message.AddRatingSuccess(result))
	return result, nil
}

func NewProductService(productRepo repo.ProductRepo, validatorSrv ValidationSrv, loggerSrv LogSrv) ProductService {
	return &productSrv{productRepo: productRepo, validatorSrv: validatorSrv, loggerSrv: loggerSrv}
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

	product.Product.DateUpdated = time.Now().Local()

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
