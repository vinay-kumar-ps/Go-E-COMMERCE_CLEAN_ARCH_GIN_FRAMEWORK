package interfaces

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
)

type PaymentUsecase interface{
	AddNewPaymentMethod(paymentMethod string)error
	RemovePaymentMethod(paymentMethodID int)error
	GetPaymentMethods()([]domain.PaymentMethod,error)
	MakePaymentRazorPay(orderID string,userID int)(models.OrderPaymentDetails,error)
	VerifyPayment(paymentID string,razorID string,orderID string)error
}