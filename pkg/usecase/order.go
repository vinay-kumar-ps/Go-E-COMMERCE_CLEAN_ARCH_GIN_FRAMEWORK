package usecase

import (
	domain "ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interfaces"
	services "ecommerce/pkg/usecase/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"
	"fmt"
)

type orderUseCase struct {
	orderRepository  interfaces.OrderRepository
	couponRepository interfaces.CouponRepository
	userUseCase      services.UserUseCase
}

func NewOrderUseCase(repo interfaces.OrderRepository, coup interfaces.CouponRepository, userUseCase services.UserUseCase) *orderUseCase {
	return &orderUseCase{
		orderRepository:  repo,
		couponRepository: coup,
		userUseCase:      userUseCase,
	}
}

func (i *orderUseCase) GetOrders(id int) ([]domain.OrderDetailsWithImages, error) {

	orders, err := i.orderRepository.GetOrders(id)
	if err != nil {
		return []domain.OrderDetailsWithImages{}, err
	}

	var result []domain.OrderDetailsWithImages
	for _, v := range orders {
		var o domain.OrderDetailsWithImages
		images, err := i.orderRepository.GetProductImagesInAOrder(int(v.ID))
		if err != nil {
			return []domain.OrderDetailsWithImages{}, err
		}

		payment, err := i.orderRepository.FindPaymentMethodOfOrder(int(v.ID))
		if err != nil {
			return []domain.OrderDetailsWithImages{}, err
		}

		o.OrderDetails = v
		o.Images = images
		o.PaymentMethod = payment

		fmt.Println("images:", images)
		fmt.Println("o.images", o.Images)

		fmt.Println("o", o)

		result = append(result, o)
	}

	return result, nil

}

func (i *orderUseCase) OrderItemsFromCart(userid int, addressid int, paymentid int, couponID int) error {

	cart, err := i.userUseCase.GetCart(userid)
	if err != nil {
		return err
	}

	var total float64
	for _, v := range cart.Data {
		total = total + v.DiscountedPrice
	}

	//finding discount if any
	coupon, err := i.couponRepository.FindCouponDetails(couponID)
	if err != nil {
		return err
	}

	totalDiscount := (total * float64(coupon.DiscountRate)) / 100

	total = total - totalDiscount

	order_id, err := i.orderRepository.OrderItems(userid, addressid, paymentid, total, coupon.Coupon)
	if err != nil {
		return err
	}

	if err := i.orderRepository.AddOrderProducts(order_id, cart.Data); err != nil {
		return err
	}

	for _, v := range cart.Data {
		if err := i.userUseCase.RemoveFromCart(cart.ID, v.ID); err != nil {
			return err
		}
	}

	return nil

}

func (i *orderUseCase) CancelOrder(id int) error {

	//the order has to be less than status delivered (pending,shipped) to be canceled
	status, err := i.orderRepository.CheckOrderStatusByID(id)
	if err != nil {
		return err
	}

	if status != "PENDING" {
		return errors.New("order cannot be canceled if you accidently booked kindly return the product")
	}

	err = i.orderRepository.CancelOrder(id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) EditOrderStatus(status string, id int) error {

	err := i.orderRepository.EditOrderStatus(status, id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) AdminOrders() (domain.AdminOrdersResponse, error) {

	var response domain.AdminOrdersResponse

	pending, err := i.orderRepository.AdminOrders("PENDING")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	shipped, err := i.orderRepository.AdminOrders("SHIPPED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	delivered, err := i.orderRepository.AdminOrders("DELIVERED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	returned, err := i.orderRepository.AdminOrders("RETURNED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	canceled, err := i.orderRepository.AdminOrders("CANCELED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	response.Canceled = canceled
	response.Pending = pending
	response.Shipped = shipped
	response.Returned = returned
	response.Delivered = delivered

	return response, nil

}

func (i *orderUseCase) ReturnOrder(id int) error {

	//should check if the order is already returned peoples will misuse this security breach
	// and will get  unlimited money into their wallet
	status, err := i.orderRepository.CheckOrderStatusByID(id)
	if err != nil {
		return err
	}

	if status == "RETURNED" {
		return errors.New("order already returned")
	}

	//should also check if the order is already returned
	//or users will also earn money by returning pending orders by opting COD

	if status != "DELIVERED" {
		return errors.New("user is trying to return an order which is still not delivered")
	}

	//make order as returned order
	if err := i.orderRepository.ReturnOrder(id); err != nil {
		return err
	}

	//find amount to be credited to user
	amount, err := i.orderRepository.FindAmountFromOrderID(id)
	if err != nil {
		return err
	}

	//find the user
	userID, err := i.orderRepository.FindUserIdFromOrderID(id)
	if err != nil {
		return err
	}
	//find if the user having a wallet
	walletID, err := i.orderRepository.FindWalletIdFromUserID(userID)
	if err != nil {
		return err
	}
	//if no wallet create new one
	if walletID == 0 {
		walletID, err = i.orderRepository.CreateNewWallet(userID)
		if err != nil {
			return err
		}
	}
	//credit the amount into users wallet
	if err := i.orderRepository.CreditToUserWallet(amount, walletID); err != nil {
		return err
	}

	return nil

}

func (i *orderUseCase) MakePaymentStatusAsPaid(id int) error {

	err := i.orderRepository.MakePaymentStatusAsPaid(id)
	if err != nil {
		return err
	}
	return nil

}

func (i *orderUseCase) GetIndividualOrderDetails(id int) (models.IndividualOrderDetails, error) {

	details, err := i.orderRepository.GetIndividualOrderDetails(id)
	if err != nil {
		return models.IndividualOrderDetails{}, err
	}

	productDetail, err := i.orderRepository.GetProductDetailsInOrder(id)
	if err != nil {
		return models.IndividualOrderDetails{}, err
	}

	details.Products = productDetail

	return details, nil
}
