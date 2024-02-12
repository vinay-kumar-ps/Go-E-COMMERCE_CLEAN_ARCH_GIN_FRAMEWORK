package handler

import (
	"ecommerce/pkg/domain"
	services "ecommerce/pkg/usecase"
	"ecommerce/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryUseCase services.CategoryUseCase

}
func NewCategoryHandler(usecase services.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		CategoryUseCase: usecase,
	}
}
// @Summary		Add Category
// @Description	Admin can add new categories for products
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			category	body	domain.Category	true	"category"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category [post]

func(cat * CategoryHandler) AddCategory(c *gin.Context){
	var category domain.Category
	if err := c.BindJSON(&category);err !=nil{
		errorRes :=response.ClientResponse(http.StatusBadRequest,"fields provided are in wrong format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	categoryResponse ,err:= cat.CategoryUseCase.AddCategory(category)
	if err !=nil{
		errorRes := response.ClientResponse(http.StatusBadRequest,"could not add the category",nil,err.Error())
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	successRes :=response.ClientResponse(http.StatusOK,"successfully added category",categoryResponse,nil)
	c.JSON(http.StatusOK,successRes)

}