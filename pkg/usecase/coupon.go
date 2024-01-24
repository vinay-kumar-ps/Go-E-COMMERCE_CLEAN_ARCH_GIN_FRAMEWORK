package usecase

import (
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interfaces"
	services "ecommerce/pkg/usecase/interfaces"
	"errors"
)

type couponUsecase struct {
	couponRepo interfaces.CouponRepository
}

// constructor function

func NewCouponUsecase(couponRepo interfaces.CouponRepository) services.CouponUsecase {
	return &couponUsecase{
		couponRepo: couponRepo,
	}
}

func (coupU *couponUsecase) AddCoupon(coupon domain.Coupon) error {
	if err := coupU.couponRepo.AddCoupon(coupon); err != nil {
		return errors.New("coupon adding failed")
	}
	return nil
}

func (coupU *couponUsecase) MakeCouponInvalid(id int) error {
	if err := coupU.couponRepo.MakeCouponInvalid(id); err != nil {
		return err
	}
	return nil
}

func (coupU *couponUsecase) GetCoupons(page, limit int) ([]domain.Coupon, error) {
	coupons, err := coupU.couponRepo.GetCoupons(page, limit)
	if err != nil {
		return []domain.Coupon{}, err
	}
	return coupons, nil
}