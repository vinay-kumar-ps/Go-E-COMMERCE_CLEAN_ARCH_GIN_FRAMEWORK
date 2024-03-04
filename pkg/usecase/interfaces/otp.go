package interfaces

import "ecommerce/pkg/utils/models"

type OtpUsecase interface {
	VerifyOTP(code models.VerifyData) (models.UserToken, error)
	SendOTP(phone string) error
}
