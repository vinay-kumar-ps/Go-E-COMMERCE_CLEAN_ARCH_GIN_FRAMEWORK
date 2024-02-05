package handler

import (
	services "ecommerce/pkg/usecase/interfaces"
	"ecommerce/pkg/utils/models"
	"ecommerce/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
type AdminHandler struct {
	adminUsecase services.AdminUsecase
}
//constructor  function

func NewAdminHandler(adminUsecase services.AdminUsecase)*AdminHandler{
	return &AdminHandler{
		adminUsecase: adminUsecase,
	}
}

     
// @Summary		Admin Login
// @Description	Login handler for admins
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.AdminLogin	true	"Admin login details"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/adminlogin [post]

func(ah *AdminHandler) LoginHandler(c *gin.Context){
	//login handler for the admin
	var adminDetails models.AdminLogin

	if err := c.BindJSON(&adminDetails);err!=nil{
		errRes :=response.ClientResponse(http.StatusBadRequest,"can't authenticate admin",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	admin,err := ah.adminUsecase.LoginHandler(adminDetails)
	if err !=nil{
		errRes :=response.ClientResponse(http.StatusBadRequest,"can't authenticate admin",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",admin.Token,3600, "/","",true,false)
	successRes :=response.ClientResponse(http.StatusOK,"Admin authenticated successfully",admin,nil)
	c.JSON(http.StatusOK,successRes)

}

// @Summary		Block User
// @Description	using this handler admins can block an user
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"user-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/users/block [post]
func(ah *AdminHandler) BlockUser(c *gin.Context) {
	id :=c.Query("id")
	err :=ah.adminUsecase.BlockUser(id)
	if err !=nil{
		errorRes:= response.ClientResponse(http.StatusBadRequest,"user could not be blocked",nil,err.Error)
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	successRes:= response.ClientResponse(http.StatusOK,"successfully blocked the user ",nil,nil)
	c.JSON(http.StatusOK,successRes)
}
// @Summary		UnBlock an existing user
// @Description	UnBlock user
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"user-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/users/unblock [POST]
func(ah *AdminHandler) UnBlockUser(c *gin.Context){
	 
	id :=c.Query("id")
	err :=ah.adminUsecase.UnblockUser(id)
	if err !=nil{
		errorRes :=response.ClientResponse(http.StatusBadRequest,"user could not be unblocked",nil,err.Error())
		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	successRes :=response.ClientResponse(http.StatusOK,"successfully unblocked the user",nil,nil)
	c.JSON(http.StatusOK,successRes)
}
// @Summary		Get Users
// @Description	Retrieve users with pagination
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			page	query		string	true	"Page number"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/users [get]

func(ah *AdminHandler)GetUsers(c *gin.Context){
	pageStr :=c.Query("page")
	page,err :=strconv.Atoi(pageStr)
	if err !=nil{
		errorRes :=response.ClientResponse(http.StatusBadRequest,"page number not in right format",nil,err.Error())

		c.JSON(http.StatusBadRequest,errorRes)
		return
	}
	users ,err :=ah.adminUsecase.GetUsers(page)
	if err !=nil{
		erroRes:= response.ClientResponse(http.StatusBadRequest,"could not retrive records",nil,err.Error())
		c.JSON(http.StatusBadRequest,erroRes)
		return
	}
}