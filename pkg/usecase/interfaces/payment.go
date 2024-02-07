package interfaces

import (
	"ecommerce/pkg/utils/models"
)
type PaymentUseCase interface {
	MakePaymentRazorPay(orderID string, userID string) (models.OrderPaymentDetails, error)
	VerifyPayment(paymentID string, razorID string, orderID string) error

	UseWallet(orderID string, userID string) (models.OrderPaymentDetails, error)
}