package usecase

import "ecommerce/pkg/repository/interfaces"

type paymentUsecase struct{
	paymentRepo interfaces.PaymentRepository
	userRepo interfaces.UserRepository
}

//constructor function
func New()