package handler

import (
	services "ecommerce/pkg/usecase/interfaces"
	models "ecommerce/pkg/utils/models"
	response "ecommerce/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	usecase services.WishlistUseCase
}

func NewWishlistHandler(use services.WishlistUseCase) *WishlistHandler {
	return &WishlistHandler{
		usecase: use,
	}
}

func (w *WishlistHandler) AddToWishlist(c *gin.Context) {

	var model models.AddToCart
	if err := c.BindJSON(&model); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	if err := w.usecase.AddToWishlist(model.UserID, model.InventoryID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add to Wishlist", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added To wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (w *WishlistHandler) RemoveFromWishlist(c *gin.Context) {
	invID, err := strconv.Atoi(c.Query("inv_id"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	UserID, ok := c.MustGet("id").(int)
	if !ok {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not find id from context", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	if err := w.usecase.RemoveFromWishlist(invID, UserID); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not remove from wishlist", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Removed product from wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (w *WishlistHandler) GetWishList(c *gin.Context) {
	pageStr := c.Query("id")
	id, err := strconv.Atoi(pageStr)

	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "id not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := w.usecase.GetWishList(id)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", products, nil)
	c.JSON(http.StatusOK, successRes)
}
