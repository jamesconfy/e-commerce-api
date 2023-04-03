package handler

import "e-commerce/internal/service"

type checkoutHandler struct{}

func NewCheckoutHandler(checkoutSrv service.CheckoutService) *checkoutHandler {
	return &checkoutHandler{}
}
