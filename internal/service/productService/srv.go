package productService

import (
	"e-commerce/internal/models/errorModels"
	"e-commerce/internal/models/productModels"
	"e-commerce/internal/service/loggerService"
	validationService "e-commerce/internal/service/validatorService"
	"e-commerce/utils"
)

type ProductService interface {
	AddProduct(userId string, req *productModels.AddProductReq) (*productModels.AddProductRes, *errorModels.ServiceError)
}

type productSrv struct {
	validatorSrv validationService.ValidationSrv
	loggerSrv    loggerService.LogSrv
}

func (p *productSrv) AddProduct(userId string, req *productModels.AddProductReq) (*productModels.AddProductRes, *errorModels.ServiceError) {
	if err := p.validatorSrv.Validate(req); err != nil {
		p.loggerSrv.Error(utils.Messages.AddProductValidationError(userId, req))
		return nil, errorModels.NewValidatingError(err)
	}

	return nil, nil
}

func NewProductService(validatorSrv validationService.ValidationSrv, loggerSrv loggerService.LogSrv) ProductService {
	return &productSrv{validatorSrv: validatorSrv, loggerSrv: loggerSrv}
}
