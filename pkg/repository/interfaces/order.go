package interfaces

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
	"time"
)

type OrderRepository interface {
	GetOrders(id int) ([]domain.Order, error)
	GetCart(userid int) ([]models.GetCart, error)
	OrderItems(userid, addressid, paymentid int, total float64, coupon string) (int, error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	CancelOrder(id int) error
	EditOrderStatus(status string, id int) error
	AdminOrders(status string) ([]domain.OrderDetails, error)

	CheckOrder(orderID string, userID int) error
	GetOrderDetail(orderID string) (domain.Order, error)

	CheckOrderStatusByID(id int) (string, error)
	ReturnOrder(id int) error
	FindAmountFromOrderID(id int) (float64, error)
	CreditToUserWallet(amount float64, walletID int) error
	FindUserIdFromOrderID(id int) (int, error)
	FindWalletIdFromUserID(userId int) (int, error)
	CreateNewWallet(userID int) (int, error)
	MakePaymentStatusAsPaid(id int) error
	GetProductImagesInAOrder(id int) ([]string, error)

	GetIndividualOrderDetails(id int) (models.IndividualOrderDetails, error)
	GetProductDetailsInOrder(id int) ([]models.ProductDetails, error)

	FindPaymentMethodOfOrder(id int) (string, error)
}