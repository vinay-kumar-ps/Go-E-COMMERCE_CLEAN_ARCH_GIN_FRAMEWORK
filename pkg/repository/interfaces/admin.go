package interfaces

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUserById(id string) (domain.User,error)
	UpdateBlockUserById(user domain.User) error
	Getusers(page, limit int) ([]models.UserDetailsAtAdmin, error)
}
