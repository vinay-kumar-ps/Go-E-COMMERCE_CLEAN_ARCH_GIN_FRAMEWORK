package interfaces

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
)

type UserRepository interface {
	CheckUserAvailability(email string) bool
	UserBlockStatus(email string) (bool, error)
	FindUserByEmail(user models.UserLogin) (models.UserResponse, error)
	SignUp(user models.UserDetails) (models.UserResponse, error)
	AddAdress(id int, address models.AddAddress, result bool) error
	GetAdresses(id int) ([]domain.Address, error)
	CheckIfFirstAddress(id int) bool
	GetUserDetails(id int) (models.UserResponse, error)
	FindUserIDByOrderID(orderID int) (int, error)
	FindUserByOrderID(orderID string) (domain.User, error)
	ChangePassword(id int, password string) error
	GetPassword(id int) (string, error)
	FindIdFromPhone(phone string) (int, error)
	EditName(id int, name string) error
	EditEmail(id int, email string) error
	EditPhone(id int, phone string) error
	EditUserame(id int, username string) error

	RemoveFromCart(id int, InventoryID int) error
	ClearCart(cartID int) error
	UpdateQuantityAdd(id, inv_id int) error
	UpdateQuantityLess(id, inv_id int) error

	GetCartID(id int) (int, error)
	GetProductsInCart(cart_id, page, limit int) ([]int, error)
	FindProductsNames(inventory_id int) (string, error)
	findCartQuantity(cart_id, inventory_id int) (int, error)
	findPrice(inventory_id int) (float64, error)
	FindCategory(inventory_id int) (int, error)
}
