package interfaces

import "ecommerce/pkg/domain"

type CouponUsecase interface{
	AddCoupon (coupon domain.Coupon)error
	MakeCouponInvalid(id int)error
	GetCoupons(page,limit int)([]domain.Coupon,error)
}