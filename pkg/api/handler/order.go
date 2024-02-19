package handler

import (
	services "ecommerce/pkg/usecase/interfaces"
	"ecommerce/pkg/utils/models"
	"ecommerce/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUseCase services.OrderUseCase
}

func NewOrderHandler(useCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		OrderUseCase: useCase,
	}

}

// @Summary		Get Orders
// @Description	user can view the details of the orders
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/orders [get]

func (i *OrderHandler) GetOrders(c *gin.Context) {
	idString := c.Query("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check your id again", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	orders, err := i.OrderUseCase.GetOrders(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrive records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfullyl got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)

} // @Summary		Order Now
// @Description	user can order the items that currently in cart
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			order body	models.Order  true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/check-out/order [post]
func (i *OrderHandler) OrderItemsFromCart(c *gin.Context) {

	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.OrderUseCase.OrderItemsFromCart(order.UserID, order.AddressID, order.PaymentMethodID, order.CouponID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not make order ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully made th order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Order Cancel
// @Description	user can cancel the orders
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			id  query  string  true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/orders [delete]

func (i *OrderHandler) CancelOrder(c *gin.Context) {
	id, err := strconv.Atoi("id")
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "conversion to integer not possible", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}
	if err := i.OrderUseCase.CancelOrder(id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "feilds provided in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully canceled the order ", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Update Order Status
// @Description	Admin can change the status of the order
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id  query  string  true	"id"
// @Param			status  query  string  true	"status"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/orders/edit/status [put]
func (i *OrderHandler) EditOrderStatus(c *gin.Context) {
	var status models.EditOrderStatus
	err := c.BindJSON(&status)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "conversion to integer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.OrderUseCase.EditOrderStatus(status.Status, status.OrderID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)

		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully edited the order status ", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Admin Orders
// @Description	Admin can view the orders according to status
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/orders [get]
func (i *OrderHandler) AdminOrders(c *gin.Context) {

	orders, err := i.OrderUseCase.AdminOrders()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Return Order
// @Description	user can return the ordered products which is already delivered and then get the amount fot that particular purchase back in their wallet
// @Tags			User
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/orders/return [put]

func (i *OrderHandler) ReturnOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		erroRes := response.ClientResponse(http.StatusBadRequest, "conversion to integer not possible", nil, err.Error())
		c.JSON(http.StatusBadRequest, erroRes)
		return
	}
	if err := i.OrderUseCase.ReturnOrder(id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "return success .the amount will be credited your wallet", nil, err.Error())
	c.JSON(http.StatusOK, successRes)
}
func (i *OrderHandler) MakePaymentStatusAsPaid(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "conversion to integer not possible", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.orderUseCase.MakePaymentStatusAsPaid(id); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully updated as paid", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *OrderHandler) GetIndividualOrderDetails(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in getting parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	details, err := i.OrderUseCase.GetIndividualOrderDetails(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not fetch the details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully fetched order details", details, nil)
	c.JSON(http.StatusOK, successRes)

}
