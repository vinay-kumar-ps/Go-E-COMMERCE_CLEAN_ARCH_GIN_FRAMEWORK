package usecase

import (
	"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/usecase/interfaces"
)

type cartusecase struct {
	cartRepo       interfaces.CartRepository
	invRepo        interfaces.InventoryRespository
	userUsecase    services.userUsecase
	paymentUsecase services.PaymentUsecase
}
