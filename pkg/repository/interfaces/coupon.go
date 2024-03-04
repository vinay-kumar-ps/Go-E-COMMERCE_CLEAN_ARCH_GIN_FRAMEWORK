package interfaces

import(
	 "ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
	)

	type CouponRepository interface{
		AddCoupon(models.Coupon)error
		MakeCouponInvalid(id int)error
		FindCouponDetails(couponId int)(domain.Coupon,error)
		GetCoupons()([]domain.Coupon,error)
		ValidateCoupon(coupon string)(bool,error)
	}