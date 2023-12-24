package interfaces

import "ecommerce/pkg/utils/models"
type OtpRepository interface{
	FindUserByMObileNumber(phone string) bool
	UserDetailsUsingPhone(phone string)(models.UserResponse,error)
}