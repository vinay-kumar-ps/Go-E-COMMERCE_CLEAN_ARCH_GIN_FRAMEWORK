package interfaces

import "ecommerce/pkg/utils/models"

type WishlistRepository interface {
	AddToWishlist(user_id, inventory_id int) error
	RemoveFromWishlist(inventory_id, UserID int) error
	GetWishList(id int) ([]models.Inventories, error)
	CheckIfTheItemIsPresentAtWishlist(userID, productID int) (bool, error)
	CheckIfTheItemIsPresentAtCart(userID, productID int) (bool, error)
}