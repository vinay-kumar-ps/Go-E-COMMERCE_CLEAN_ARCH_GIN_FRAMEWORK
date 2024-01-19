package usecase

import (
	"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
)

type adminUsecase struct {
	adminRepository interfaces.AdminRepository
}

// constructor function
func NewAdminUsecase(adRepo interfaces.AdminRepository) services.AdminUsecase {
	return &adminUsecase{
		adminRepository: adRepo,
	}
}
func(au *adminUsecase) LoginHandler(adminDetails models.AdminLogin)(models.AdminToken,error){
	
}

