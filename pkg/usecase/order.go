package usecase

import (
	interfaces "ecommerce/pkg/repository/interfaces"
    services"ecommerce/pkg/usecase/interfaces"  
)

type OrderUsecase struct{
	orderRepo  interfaces.OrderRepository
	userUsecase  services.UserUsecase
	walletRepo interfaces.WalletRepository
	couponRepo  interfaces.CouponRepository

}


