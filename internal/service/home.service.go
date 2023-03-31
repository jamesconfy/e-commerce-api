package service

type HomeService interface {
	CreateHome() (string, error)
}

type homeSrv struct {
	loggerSrv LogSrv
}

func (h *homeSrv) CreateHome() (string, error) {
	// h.loggerSrv.Info("You have gotten to home page")
	return "Home", nil
}

func NewHomeService(loggerSrv LogSrv) HomeService {
	return &homeSrv{loggerSrv: loggerSrv}
}
