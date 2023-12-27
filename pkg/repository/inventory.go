package repository

import (
	"ecommerce/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type inventoryRepository struct{
	DB *gorm.DB

}
//constructor function

func NewInventoryRepository(DB *gorm.DB)interfaces.InventoryRepository{

	return &inventoryRepository{
		DB: DB,
	}
}