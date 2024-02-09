package handler

import (
services"ecommerce/pkg/usecase/interfaces"
	"ecommerce/pkg/utils/models"
	"ecommerce/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	usecase services.CartUsecase
}

func NewCartHandler(usecase services.CartUsecase) *CartHandler {
	return &CartHandler{
		usecase: usecase,
	}
}

// @Summary		Add To Cart
// @Description	Add products to carts  for the purchase
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			cart	body	models.AddToCart	true	"Add To Cart"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/home/add-to-cart [post]

func (i *CartHandler) AddToCart(c *gin.Context) {
	var model models.AddToCart
	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := i.usecase.AddToCart(model.UserID, model.InventoryID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not add the product to cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully added to cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
func (i *CartHandler) CheckOut(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "user_id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	products, err := i.usecase.CheckOut(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not open checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Succesfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)

}
