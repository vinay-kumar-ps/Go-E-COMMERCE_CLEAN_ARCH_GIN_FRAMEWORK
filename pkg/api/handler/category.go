package handler

import (
	"ecommerce/pkg/domain"
	services "ecommerce/pkg/usecase"
	"ecommerce/pkg/utils/models"
	"ecommerce/pkg/utils/response"
	"net/http"
	"strconv"

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

func(Cat * CategoryHandler) AddCategory(c *gin.Context){
	var category domain.Category
	if err := c.BindJSON(&category);err !=nil{
		errorRes :=response.ClientResponse(http.StatusBadRequest,"fields provided are in wrong format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	categoryResponse ,err:= Cat.CategoryUseCase.AddCategory(category)
	if err !=nil{
		errorRes := response.ClientResponse(http.StatusBadRequest,"could not add the category",nil,err.Error())
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	successRes :=response.ClientResponse(http.StatusOK,"successfully added category",categoryResponse,nil)
	c.JSON(http.StatusOK,successRes)

}
// @Summary		Update Category
// @Description	Admin can update name of a category into new name
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			set_new_name	body	models.SetNewName	true	"set new name"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category [put]

func (Cat *CategoryHandler) UpdateCategory(c *gin.Context){
	var p models.SetNewName
	if err :=c.BindJSON(&p);err !=nil{
		erroRes:=response.ClientResponse(http.StatusBadRequest,"fields provided in wrong format",nil,err.Error())
		c.JSON(http.StatusBadRequest,erroRes)
		return
	}
	a,err :=Cat.CategoryUseCase.UpdateCategory(p.Current,p.New)
	if err !=nil{
		errorRes :=response.ClientResponse(http.StatusBadRequest,"could not update the category ",nil,err.Error())
	    c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	succesRes :=response.ClientResponse(http.StatusOK,"successfulyy renamed the category",a,nil)
	c.JSON(http.StatusOK,succesRes)
}
// @Summary		Delete Category
// @Description	Admin can delete a category
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category [delete]
func(Cat *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryID :=c.Query("id")
	err :=Cat.CategoryUseCase.DeleteCategory(categoryID)
	if err !=nil{
		errorRes:=response.ClientResponse(http.StatusBadRequest,"fields provided arein wrong format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	succesRes:=response.ClientResponse(http.StatusOK,"successfully deleted the category ",nil,nil)
	c.JSON(http.StatusOK,succesRes)

}
func(Cat *CategoryHandler)GetCategory(c *gin.Context){
	categories,err :=Cat.CategoryUseCase.GetCategories()
	if err !=nil{
		errorRes :=response.ClientResponse(http.StatusBadRequest,"fields rovided are in wrong format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	succesRes :=response.ClientResponse(http.StatusOK,"successfully got all categories ",categories,nil)
	c.JSON(http.StatusOK,succesRes)
}
func(Cat *CategoryHandler) GetProductsDetailsInCategory(c *gin.Context){
	id ,err:=strconv.Atoi(c.Query("id"))
	if err !=nil{
		errorRes :=response.ClientResponse(http.StatusBadRequest,"feilds provided are in wrong format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	products,err :=Cat.CategoryUseCase.GetProductsDetailsInCategory(id)
	if err !=nil{
		erroRes :=response.ClientResponse(http.StatusInternalServerError,"error in fetcing data",nil,err.Error())
		c.JSON(http.StatusBadRequest,erroRes)
		return
	}
	successRes:=response.ClientResponse(http.StatusOK,"successfully got all categories",products,nil)
	c.JSON(http.StatusOK,successRes)	
}
func (Cat *CategoryHandler)GetBannerForUsers (c *gin.Context){

	banners,err :=Cat.CategoryUseCase.GetBannerForUsers()
	if err !=nil{
		errorRes :=response.ClientResponse(http.StatusInternalServerError,"error in fetching data",nil,err.error())
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	successRes :=response.ClientResponse(http.StatusOK,"successfully got all banners",banners,nil)
	c.JSON(http.StatusOK,successRes)
}