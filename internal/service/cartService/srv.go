package cartService

import (
	"e-commerce/internal/service/loggerService"
	validationService "e-commerce/internal/service/validatorService"
)

type CartService interface {
}

type cartSrv struct {
	loggerSrv    loggerService.LogSrv
	validatorSrv validationService.ValidationSrv
}

func NewCartService(loggerSrv loggerService.LogSrv, validatorSrv validationService.ValidationSrv) CartService {
	return &cartSrv{loggerSrv: loggerSrv, validatorSrv: validatorSrv}
}
