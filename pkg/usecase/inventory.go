package usecase

import (
	"ecommerce/pkg/repository/interfaces"
	services "ecommerce/pkg/usecase/interfaces"
)

type InventoryUsecase struct {
	invRepo interfaces.InventoryRespository
}

//contructor function

func NewInventoryUsecase(invRepo interfaces.InventoryRespository) services.InventoryUsecase {
	return InventoryUsecase{
		invRepo: invRepo,
	}
}
func 
