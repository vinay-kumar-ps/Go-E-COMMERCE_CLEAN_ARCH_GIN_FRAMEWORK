package handler

import(
	services "ecommerce/pkg/usecase/interfaces"
)

type InventoryHandler struct{
	InventoryUseCase services.InventoryUseCase
}
func NewInventoryHandler (usecase services.InventoryUseCase)*InventoryHandler{
	return &InventoryHandler{
		InventoryUseCase: usecase,
	}
}
