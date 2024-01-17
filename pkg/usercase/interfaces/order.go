package interfaces

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/utils/models"
)

type OrderUsecase interface{
	GetOrders(id, page ,limit int) ([]domain.Order,error)
	OrderItemsFromCart(userId int, order models.Order,coupon string)(string,error)
	CancelOrder(id ,orderId int)error
	EditOrderStatus(status string,id int)error
	MarkAsPaid(orderId int )error
	AdminOrders(page,limit int, status string)([]domain.OrderDetails,error)
	DailyOrders()(domain.SalesReport,error)
	WeeklyOrders()(domain.SalesReport,error)
	MonthlyOrders()(domain.SalesReport,error)
	AnnualOrders()(domain.SalesReport,error)
	CustomDateOrders(dates models.CustomDates)(domain.SalesReport,error)
	ReturnOrder(id int )error
}