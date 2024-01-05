package repository

import (
	"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"

	"gorm.io/gorm"
)

type otpRepository struct{
	DB *gorm.DB
}
func NewOtpRepository (DB *gorm.DB) interfaces.OtpRepository {
	return &otpRepository{
		DB: DB,
	}
}
func (otr *otpRepository) FindUserByMobileNumber(phone string)bool{
	var count int
	err := otr.DB.Raw("SELECT COUNT (*) users WHERE phone =?",phone).Scan(&count).Error
	if err != nil{
		return false
	}
	return count >0
}
func (otr *otpRepository) UserDetailsUsingPhone(phone string) (models.UserResponse,error){
	var userDetails  models.UserResponse 
	err :=otr.DB.Raw("SELECT * FROM users WHERE phone=?",phone).Scan(&userDetails).Error
	
	if err != nil{
		return models.UserResponse{},err

	}
	return userDetails,nil

}