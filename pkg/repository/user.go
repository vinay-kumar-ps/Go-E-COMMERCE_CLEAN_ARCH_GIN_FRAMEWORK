package repository

import (
	"ecommerce/pkg/repository/interfaces"

	"gorm.io/gorm"
)

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository{
	return &userRepository[
		DB: DB,
	]

}
 func (ur *userRepository) CheckUserAvailability(email string )bool{

var userCount int

err := ur.DB.Raw("SELECT COUNT (*) FROM users WHERE email=? ",email).Scan(&userCount).Error

	
}