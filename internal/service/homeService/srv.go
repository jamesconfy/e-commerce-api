package homeService

type HomeService interface {
	CreateHome() (string, error)
}

type homeSrv struct{}

func (h *homeSrv) CreateHome() (string, error) {
	return "Home", nil
}

func NewHomeSrv() HomeService {
	return &homeSrv{}
}
