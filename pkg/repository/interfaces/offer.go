package interfaces

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
)

type OfferRepository interface {
	AddNewOffer(models.CreateOffer) error
	makeOfferExpired(categoryId int) error
	FindDiscountPercentage(categoryId int) (int, error)
	GetOffers(page, limit int) ([]domain.Offer, error)
}
