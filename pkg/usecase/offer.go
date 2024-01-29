package usecase

import (
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interfaces"
	services "ecommerce/pkg/usecase/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"
)

type offerUsecse struct{
	offerRepo interfaces.OfferRepository


}
//constructor function

func NewOfferUsecase(offerRepo interfaces.OfferRepository) services.OfferUsecase{
	return &offerUsecse{
		offerRepo:offerRepo,
	}
}
func (offU *offerUsecse)AddNewOffer(model models.CreateOffer)error{
	if err := offU.offerRepo.AddNewOffer(model);err!=nil{
		return errors.New("adding offer failed")
	}
	return nil
}
func(offU *offerUsecse)MakeOfferExpire(catID int )error{
	if err :=offU.offerRepo.MakeOfferExpired(catID);err!=nil{
		return err
	}
	return nil
}

func (offU *offerUsecse) GetOffers(page, limit int) ([]domain.Offer, error) {
	offers, err := offU.offerRepo.GetOffers(page, limit)
	if err != nil {
		return []domain.Offer{}, err
	}
	return offers, nil
}
