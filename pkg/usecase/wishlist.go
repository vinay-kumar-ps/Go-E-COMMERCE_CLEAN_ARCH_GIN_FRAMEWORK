package usecase

import (
	interfaces "ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"
)

type wishlistUseCase struct {
	repository interfaces.WishlistRepository
	offerRepo  interfaces.OfferRepository
}

func NewWishlistUseCase(repo interfaces.WishlistRepository, offer interfaces.OfferRepository) *wishlistUseCase {
	return &wishlistUseCase{
		repository: repo,
		offerRepo:  offer,
	}
}

func (w *wishlistUseCase) AddToWishlist(userID, inventoryID int) error {

	exists, err := w.repository.CheckIfTheItemIsPresentAtWishlist(userID, inventoryID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("item already exists in wishlist")
	}

	if err := w.repository.AddToWishlist(userID, inventoryID); err != nil {
		return errors.New("could not add to wishlist")
	}

	return nil
}

func (w *wishlistUseCase) RemoveFromWishlist(inventoryID, UserID int) error {

	if err := w.repository.RemoveFromWishlist(inventoryID, UserID); err != nil {
		return errors.New("could not remove from wishlist")
	}

	return nil
}

func (w *wishlistUseCase) GetWishList(id int) ([]models.Inventories, error) {

	productDetails, err := w.repository.GetWishList(id)
	if err != nil {
		return []models.Inventories{}, err
	}

	//loop inside products and then calculate discounted price of each then return
	for j := range productDetails {
		discount_percentage, err := w.offerRepo.FindDiscountPercentage(productDetails[j].CategoryID)
		if err != nil {
			return []models.Inventories{}, errors.New("there was some error in finding the discounted prices")
		}
		var discount float64

		if discount_percentage > 0 {
			discount = (productDetails[j].Price * float64(discount_percentage)) / 100
		}

		productDetails[j].DiscountedPrice = productDetails[j].Price - discount

		productDetails[j].IfPresentAtWishlist, err = w.repository.CheckIfTheItemIsPresentAtWishlist(id, int(productDetails[j].ID))
		if err != nil {
			return []models.Inventories{}, errors.New("error while checking ")
		}

		productDetails[j].IfPresentAtCart, err = w.repository.CheckIfTheItemIsPresentAtCart(id, int(productDetails[j].ID))
		if err != nil {
			return []models.Inventories{}, errors.New("error while checking ")
		}

	}

	return productDetails, nil

}
