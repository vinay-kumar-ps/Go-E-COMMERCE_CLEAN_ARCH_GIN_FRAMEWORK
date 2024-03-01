package domain

import "ecommerce/pkg/utils/models"

//admin represents an administrative user in the system.

type Admin struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"validate:required"`
	UserName string `json:"email" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
}
type TokenAdmin struct {
	Admin        models.AdminDetailsResponse
	AccessToken  string
	RefreshToken string
}
