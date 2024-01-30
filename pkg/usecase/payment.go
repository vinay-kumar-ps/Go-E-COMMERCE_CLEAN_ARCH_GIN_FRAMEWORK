package usecase

import (
	"ecommerce/pkg/domain"
	interfaces "ecommerce/pkg/repository/interfaces"
	 services "ecommerce/pkg/usecase/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"

	"github.com/razorpay/razorpay-go"
)

type paymentUsecase struct {
	paymentRepo interfaces.PaymentRepository
	userRepo    interfaces.UserRepository
}

//constructor function
func NewPaymentUsecase(paymentRepo interfaces.PaymentRepository, useRepo interfaces.UserRepository) services.PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
		userRepo:    useRepo,
	}
}
func (payU *paymentUsecase) AddNewPaymentMethod(paymentMethod string) error {
	if paymentMethod == "" {
		return errors.New("enter payment method")
	}
	if err := payU.paymentRepo.AddNewPaymentMethod(paymentMethod); err != nil {
		return err
	}
	return nil
}
func (payU *paymentUsecase) RemovePaymentMethod(paymentMethodID int) error {
	if paymentMethodID == 0 {
		return errors.New("enter method id")

	}
	if err := payU.paymentRepo.RemovePaymentMethod(paymentMethodID); err != nil {
		return err
	}
	return nil
}
func (payU *paymentUsecase) GetPaymentMethods() ([]domain.PaymentMethod, error) {
	paymentMethods, err := payU.paymentRepo.GetPaymentMethods()
	if err != nil {
		return []domain.PaymentMethod{}, err
	}
	return paymentMethods, nil
}

func (payU *paymentUsecase) MakePaymentRazorPay(orderID string, userID int) (models.OrderPaymentDetails, error) {


var  orderDetails models.OrderPaymentDetails


	//Get order id
	orderId, err := strconv.Atoi(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.OrderID = orderId
	orderDetails.UserID = userID

	//Get username
	username, err := payU.paymentRepo.FindUsername(userID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.Username = username

	//Get total

	total, err := payU.paymentRepo.FindPrice(orderId)
	if err != nil {
		return models.OrderPaymentDetails{}, err

	}
	orderDetails.FinalPrice = total

	//need to add key and secret

	client := razorpay.NewClient("key", "secret")

	data := map[string]interface{}{
		"amount ":   int(orderDetails.FinalPrice) * 100,
		"currency ": "INR",
		"receipt ":  "some receipt id",
	}
	fmt.Println("razorpay::91", orderDetails, data)

	body, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Println(err)
		return models.OrderPaymentDetails{}, err
	}
	razorpayOrderId := body["id"].(int)
	orderDetails.RazorID = razorpayOrderId

	fmt.Println("razorpay::100", orderDetails)
	return orderDetails, nil

}

func(payU *paymentUsecase) VerifyPayment(paymentID string,razorID string , orderID string)error{
	if err :=payU.paymentRepo.UpdatePaymentDetails(orderID,paymentID,razorID);err !=nil{
		return err

		//clear cart

		orderIdint,err :=strconv.Atoi(orderID)
		if err !=nil{
			return err
		}
		userId ,err :=payU.userRepo.FindUserByOrderID(orderIdint)
		if err !=nil{
			return err
		}
		cartId,err :=payU.userRepo.GetCartID(userId)
		if err !=nil{
			return err
		}
		if err :=payU.userRepo.ClearCart(cartId);err !=nil{

		}
		return nil
	}
}
