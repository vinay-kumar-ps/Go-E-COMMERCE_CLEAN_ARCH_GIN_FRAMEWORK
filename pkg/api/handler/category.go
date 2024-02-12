package handler

import (
	"ecommerce/pkg/domain"
	services "ecommerce/pkg/usecase"

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
	if err
}