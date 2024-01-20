package helper

import (
	"ecommerce/pkg/utils/models"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserId (c *gin.Context)(int ,error){
	var key models.UserKey ="user_id"
	val  :=c.Request.Context().Value(key)
	
	//check if the value is not nil 
	if val ==nil{
		return 0,errors.New("user id not found in context")

	}
	//using type assertion to each the type of val is models.userkey

	userkey ,ok 
}