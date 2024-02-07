package usecase

import (
	domain "ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
)

type offerUseCase struct {
	repository interfaces.OfferRepository
}

func NewOfferUseCase(repo interfaces.OfferRepository) *offerUseCase {
	return &offerUseCase{
		repository: repo,
	}
}

func (off *offerUseCase) AddNewOffer(model models.OfferMaking) error {
	if err := off.repository.AddNewOffer(model); err != nil {
		return err
	}

	return nil
}

func (off *offerUseCase) MakeOfferExpire(id int) error {
	if err := off.repository.MakeOfferExpire(id); err != nil {
		return err
	}

	return nil
}

func (o *offerUseCase) GetOffers() ([]domain.Offer, error) {

	offers, err := o.repository.GetOffers()
	if err != nil {
		return []domain.Offer{}, err
	}
	return offers, nil

}
