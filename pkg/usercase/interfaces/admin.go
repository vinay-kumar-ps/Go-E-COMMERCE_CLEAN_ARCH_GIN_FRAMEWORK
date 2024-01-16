package interfaces

import "ecommerce/pkg/utils/models"

type AdminUsecase interface{
	LoginHandler(adminDetails models.AdminLogin) (models.AdminToken, error)
	BlockUser(id string)error
	UnblockUser(id string)error
	GetUsers(page ,limit int) ([]models.UserDetailsAtAdmin,error)
}