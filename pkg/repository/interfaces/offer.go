package interfaces

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
)


type OfferRepository interface {
	AddNewOffer(models.CreateOffer) error
	MakeOfferExpired(categorytId int) error
	FindDiscountPercentage(categorytId int) (int, error)
	GetOffers(page, limit int) ([]domain.Offer, error)
}