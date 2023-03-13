package homeService

import loggerservice "e-commerce/internal/service/loggerService"

type HomeService interface {
	CreateHome() (string, error)
}

type homeSrv struct {
	loggerSrv loggerservice.LogSrv
}

func (h *homeSrv) CreateHome() (string, error) {
	h.loggerSrv.Info("You have gotten to home page")
	return "Home", nil
}

func NewHomeSrv(loggerSrv loggerservice.LogSrv) HomeService {
	return &homeSrv{loggerSrv: loggerSrv}
}
