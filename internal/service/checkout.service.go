package service

type CheckoutService interface{}

type checkoutSrv struct{}

func NewCheckoutService() CheckoutService {
	return &checkoutSrv{}
}
