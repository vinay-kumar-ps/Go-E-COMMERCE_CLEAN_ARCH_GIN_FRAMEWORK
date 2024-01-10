package repository

import (
	"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"

	"golang.org/x/mod/module"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// constructor funciton

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (ur *userRepository) CheckUserAvailability(email string) bool {
	var userCount int

	err := ur.DB.Raw("SELECT COUNT(*) FROM users WHERE email=?", email).Scan(&userCount).Error
	if err != nil {
		return false
	}
	// if count greater than 0, user already exist
	return userCount > 0
}

func (ur *userRepository) UserBlockStatus (email string) (bool ,error){
	var permission bool
	err :=ur.DB.Raw("SELECT permission FROM users WHERE email =?",email).Scan(&permission).Error
	if err != nil{
		return false,err
	}
	return permission,nil
}
func (ur *userRepository) FindUserByEmail (user models.UserLogin) (models.UserResponse,error){
	var UserResponse models.UserResponse

	err := ur .DB.Raw("SELECT *FROM users WHERE email=? AND permission=true",user.Email).Scan(&UserResponse).Error
	if err !=nil{
		return models.UserResponse{},errors.New("no user found ")

	}
	return UserResponse,nil
}
func(ur *userRepository)FindUserIDByOrderID(orderID int )(int, error) {

	var userId int
	err := ur.DB.Raw("SELECT user_id FROm orders WHERE order_id=?",orderID).Scan(&userId).Error
	 if err!=nil{
		return 0,errors.New("user Id not found")

	 }
	 return userId ,nil
}
func (ur *userRepository)SignUp(user  models.UserDetails)(models.UserResponse,error){
	var userResponse models.UserResponse
	err := ur.DB.Exec("INSERT INTO users (name ,email,username,phone,password)VALUES(?,?,?,?,?) RETURNING id,name,email,phone",user.Name,user.Email,user.Username,user.Phone,user.Password).Scan(&userResponse).Error
	if err!=nil{
		return models.UserResponse{},err
	}
	return userResponse,nil
}
func (ur *userRepository)AddAddress( id int ,address models.AddAddress,result bool)error{
	
}