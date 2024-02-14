package handler

import(
	services "ecommerce/pkg/usecase/interfaces"
	)

type CouponHandler struct {
	uscase services.CouponUsecase
}
func NewCouponHandler(use services.CouponUsecase) *CouponHandler {
	return &CouponHandler{
		uscase: use,
	}
}

