package repository

import (
	"ecommerce/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type walletRepository struct {
	DB *gorm.DB
}
// 
func NewWalletRepository ( DB *gorm.DB) interfaces.WalletRepository{
	return &walletRepository{
		DB: DB,
	}
}
