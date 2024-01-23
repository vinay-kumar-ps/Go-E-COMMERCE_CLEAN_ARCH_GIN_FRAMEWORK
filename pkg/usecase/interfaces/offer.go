package interfaces

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
)

 type OfferUsecase interface{
	AddNewOffer(model models.CreateOffer)error
	MakeOfferExpire(cartID int )error
	GetOffers(page ,limit int)([]domain.Offer,error)
 }