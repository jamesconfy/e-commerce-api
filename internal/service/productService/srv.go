package productService

import (
	"e-commerce/internal/Repository/productRepo"
	"e-commerce/internal/models/errorModels"
	"e-commerce/internal/models/productModels"
	"e-commerce/internal/service/loggerService"
	"e-commerce/internal/service/timeService"
	validationService "e-commerce/internal/service/validatorService"
	"e-commerce/utils"
	"fmt"

	"github.com/google/uuid"
)

type ProductService interface {
	AddProduct(req *productModels.AddProductReq) (*productModels.AddProductRes, *errorModels.ServiceError)
	GetProducts(page int) ([]*productModels.GetProduct, *errorModels.ServiceError)
	GetProduct(productId string) (*productModels.GetProduct, *errorModels.ServiceError)
	DeleteProduct(productId string) (*productModels.DeleteProduct, *errorModels.ServiceError)
	AddRating(req *productModels.AddRatingsReq) (*productModels.AddRatingsRes, *errorModels.ServiceError)
	VerifyUserRatings(userId, productId string) *errorModels.ServiceError
}

type productSrv struct {
	productRepo  productRepo.ProductRepo
	validatorSrv validationService.ValidationSrv
	loggerSrv    loggerService.LogSrv
	timeSrv      timeService.TimeService
}

func (p *productSrv) AddProduct(req *productModels.AddProductReq) (*productModels.AddProductRes, *errorModels.ServiceError) {
	if err := p.validatorSrv.Validate(req); err != nil {
		p.loggerSrv.Error(utils.Messages.AddProductValidationError(req))
		return nil, errorModels.NewValidatingError(err)
	}
	// fmt.Printf("Type of provided price: %T\n", req.Price)
	// fmt.Printf("User Id: %s", req.UserId)

	req.DateCreated = p.timeSrv.CurrentTime()
	req.DateUpdated = p.timeSrv.CurrentTime()
	req.ProductId = uuid.New().String()

	err := p.productRepo.AddProduct(req)
	if err != nil {
		p.loggerSrv.Error(utils.Messages.AddProductRepoError(req, err))
		return nil, errorModels.NewCustomServiceError("Error when creating new product", err)
	}

	data := &productModels.AddProductRes{
		ProductId:   req.ProductId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Image:       req.Image,
	}

	p.loggerSrv.Info(utils.Messages.AddProductSuccess(req))
	return data, nil
}

func (p *productSrv) GetProducts(page int) ([]*productModels.GetProduct, *errorModels.ServiceError) {
	products, err := p.productRepo.GetProducts(page)
	if err != nil {
		p.loggerSrv.Error(utils.Messages.GetAllProductsRepoError(err))
		return nil, errorModels.NewCustomServiceError("Error when getting products", err)
	}

	return products, nil
}

func (p *productSrv) GetProduct(productId string) (*productModels.GetProduct, *errorModels.ServiceError) {
	product, err := p.productRepo.GetProduct(productId)
	if err != nil {
		p.loggerSrv.Error(utils.Messages.GetProductRepoError(productId, err))
		return nil, errorModels.NewCustomServiceError("Error when getting product", err)
	}

	return product, nil
}

func (p *productSrv) DeleteProduct(productId string) (*productModels.DeleteProduct, *errorModels.ServiceError) {
	product, err := p.productRepo.DeleteProduct(productId)
	if err != nil {
		p.loggerSrv.Error(utils.Messages.DeleteProductRepoError(productId, err))
		return nil, errorModels.NewCustomServiceError("Error when deleting product", err)
	}

	return product, nil
}

func (p *productSrv) AddRating(req *productModels.AddRatingsReq) (*productModels.AddRatingsRes, *errorModels.ServiceError) {
	if err := p.validatorSrv.Validate(req); err != nil {
		return nil, errorModels.NewValidatingError(err)
	}

	req.RatingId = uuid.New().String()
	req.DateCreated = p.timeSrv.CurrentTime()
	req.DateUpdated = p.timeSrv.CurrentTime()

	err := p.productRepo.AddRating(req)
	if err != nil {
		return nil, nil
	}

	data := &productModels.AddRatingsRes{
		RatingId:    req.RatingId,
		Rating:      req.Rating,
		ProductId:   req.ProductId,
		UserId:      req.UserId,
		DateCreated: req.DateCreated,
		DateUpdated: req.DateUpdated,
	}

	return data, nil
}

func (p *productSrv) VerifyUserRatings(userId, productId string) *errorModels.ServiceError {
	err := p.productRepo.VerifyUserRatings(userId, productId)
	if err != nil {
		p.loggerSrv.Error(fmt.Sprintf("User tried to re-rate a product || UserId: %s || ProductId: %s", userId, productId))
		return errorModels.NewCustomServiceError("This user has already rated this product", err)
	}

	p.loggerSrv.Error(fmt.Sprintf("Product rated successfully || UserId: %s || ProductId: %s", userId, productId))
	return nil
}

func NewProductService(productRepo productRepo.ProductRepo, validatorSrv validationService.ValidationSrv, loggerSrv loggerService.LogSrv, timeSrv timeService.TimeService) ProductService {
	return &productSrv{productRepo: productRepo, validatorSrv: validatorSrv, loggerSrv: loggerSrv, timeSrv: timeSrv}
}
