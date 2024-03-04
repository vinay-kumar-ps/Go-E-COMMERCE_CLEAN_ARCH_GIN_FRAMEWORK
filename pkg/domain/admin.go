package domain

import "ecommerce/pkg/utils/models"

//admin represents an administrative user in the system.


type Admin struct {
	ID       int    `json:"id" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"validate:required"`
	UserName string `json:"email" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
}
type AdminToken struct {
	Admin        models.AdminDetailsResponse
	Token        string
	RefreshToken string
}
