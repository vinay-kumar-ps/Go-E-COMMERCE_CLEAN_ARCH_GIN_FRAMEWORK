package handler

import (
	services "ecommerce/pkg/usecase/interfaces"
	"ecommerce/pkg/utils/models"
	"ecommerce/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type CartHandler struct {
	cartUsecase services.CartUseCase
}

// Constructor function
func NewCartHandler(usecase services.CartUseCase) *CartHandler {
	return &CartHandler{
		cartUsecase: usecase,
	}
}

// @Summary		Add To Cart
// @Description	Add products to carts  for the purchase
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			inventory	query	string	true	"inventory ID"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/home/add-to-cart [post]
func (ch *CartHandler) AddtoCart(c *gin.Context) {
	// userId, err := helper.GetUserId(c)
	var cart models.AddToCart
	if err := c.BindJSON(&cart); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := ch.cartUsecase.AddToCart(cart.UserID, cart.InventoryID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added To cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Checkout section
// @Description	Add products to carts  for the purchase
// @Tags			User
// @Produce		    json
// @Security		Bearer
// @Param          id    query   string   true  "id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/check-out [get]
func (ch *CartHandler) CheckOut(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := ch.cartUsecase.CheckOut(userId)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}