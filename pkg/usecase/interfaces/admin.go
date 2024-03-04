package interfaces

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
)


type AdminUsecase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.AdminToken, error)
	BlockUser(id string) error
	UnblockUser(id string) error
	GetUsers(page, limit int) ([]models.UserDetailsAtAdmin, error)
}