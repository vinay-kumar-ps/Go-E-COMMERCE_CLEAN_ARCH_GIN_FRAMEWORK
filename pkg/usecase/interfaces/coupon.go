package interfaces

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
)

type CouponUsecase interface {
	Addcoupon(coupon models.Coupon) error
	MakeCouponInvalid(id int) error
	GetCoupons() ([]domain.Coupon, error)
}
