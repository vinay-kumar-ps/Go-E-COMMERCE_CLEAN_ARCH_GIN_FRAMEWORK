package interfaces

import "ecommerce/pkg/utils/models"

type WishListRepository interface {
	GetWishlist(id int) ([]models.GetWishlist, error)
	GetWishlistId(user_id int) (int, error)
	CreateNewWishlist(user_id int) (int, error)
	AddWishlistItem(wishlistId, inventoryId int) error
	GetProductInWishlist(wishlistId int) ([]int, error)
	FindProductNames(inventory_id int) (string, error)
	FindPrice(inventory_id int) (float64, error)
	FindCategory(inventory_id int) (string, error)
	RemoveFromWishlist(wishlisId, inventoryId int) error
}
