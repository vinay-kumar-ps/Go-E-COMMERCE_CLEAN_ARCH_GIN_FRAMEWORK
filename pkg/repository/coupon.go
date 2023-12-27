package repository

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/repository/interfaces"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type couponRepository struct {
	DB *gorm.DB
}

// constructor function
func NewCouponRepository(DB *gorm.DB) interfaces.CouponRepository {
	return &couponRepository{
		DB: DB,
	}
}

func (couprep *couponRepository) AddCoupon(coupon domain.Coupon) error {
	err := couprep.DB.Exec("INSERT INTO coupons(name,discount_rate,valid)VALUES ($1,$2,$3)", coupon.Name, coupon.DiscountRate, coupon.Valid).Error
	if err != nil {
		return err
	}
	return nil
}

func (couprep *couponRepository) MakeCouponInvalid(id int) error {
	err := couprep.DB.Exec("UPDATE coupons SET valid=false WHERE id=$1", id).Error
	if err != nil {
		return err
	}
	return nil
}
func (couprep *couponRepository) FindCouponDiscount(coupon string) int {
	var DiscountRate int

	err := couprep.DB.Raw("SELECT discount_rate FROM coupons WHERE name=?", coupon).Scan(&DiscountRate).Error
	if err != nil {
		return 0

	}
	fmt.Println("dicount rate: ", DiscountRate)
	return DiscountRate
}
func (couprep *couponRepository) GetCoupons(page, limit int) ([]domain.Coupon, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var couponResponse []domain.Coupon
	err := couprep.DB.Raw("SELECT id,name,discount_rate,valid FROM coupons limit ? offset ?", limit, offset).Scan(&couponResponse).Error
	if err != nil {
		return []domain.Coupon{}, err
	}
	return couponResponse, nil
}
func (couponrep *couponRepository) ValidateCoupon(coupon string) (bool, error) {
	count := 0
	err := couponrep.DB.Raw("SELECT COUNT (id)FROM coupons WHERE name=?", coupon).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, errors.New("not a valid coupon")
	}
	valid := true
	err = couponrep.DB.Raw("SELECT valid FROM coupons WHERE name =?", coupon).Scan(&valid).Error
	if err != nil {
		return false, err
	}
	return true, nil

}
