package usecase

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/helper"
	interfaces "ecommerce/pkg/repository/interfaces"
	services "ecommerce/pkg/usecase/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"
	"hash"

	"golang.org/x/crypto/bcrypt"
)
type userUsecase struct{
	userRepo interfaces.UserRepository
	offerRepo interfaces.OfferRepository
	walletRepo interfaces.WalletRepository

}

func NewUserUsecase(userRepo interfaces.UserRepository,offerRepo ,walletRepo interfaces.WalletRepository)services.UserUsecase{
	return&userUsecase{
		userRepo: userRepo,
		offerRepo: offerRepo,
       walletRepo: walletRepo,
	}
}
func(usrU *userUsecase) Login(user models.UserLogin)(models.UserToken, error){
	//check the user already exist or not

	ok := usrU.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.UserToken{}, errors.New("user not exist")
	}
	//Get the user details in order to check password 
	userDetails ,err :=usrU.userRepo.FindUserByEmail(user)
	if err !=nil{
		return models.UserToken{},err
	}
	//check the password
	err =bcrypt.CompareHashAndPassword([]byte(userDetails.Password),[]byte(user.Password))
	if err !=nil{
		return models.UserToken{},errors.New("password is incorrrect")
	}
	//generate token 
	tokenString ,err :=helper.GenerateUserToken(userDetails)
	if err !=nil{
		return models.UserToken{},errors.New("couldn't create token for user")
	}
	return models.UserToken{
		Username: userDetails.Username,
		Token: tokenString,
	},nil
}
func(usrU *userUsecase)signUp(user models.UserDetails)(models.UserToken,error){
	//check the user exist or not ,if exist show the error (its a signup function)
	userExist :=usrU.userRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.UserToken{},errors.New("user already exist please sign in ")

	}
	if user .Password!=user.ConfirmPassword{
		return models.UserToken{},errors.New("password does't match")
	}
	//hash the password
	hashedPass,err :=helper.PasswordHashing(user.Password)
	if err !=nil{
		return models.UserToken{},err
	}
	user.Password=hashedPass
	//insert the user into database
	userData,err :=usrU.userRepo.SignUp(user)
	if err !=nil{
		return models.UserToken{},err
	}
	//create jwt token for user
	tokenString ,err:= helper.GenerateUserToken(userData)
	if err !=nil{
		return models.UserToken{},errors.New("could't create token for user due to some internal error")
		
	}
	return models.UserToken{
		Username: user.Username,
		Token: tokenString,

	},nil
	
}
func (usrU *userUsecase) AddAddress(id int ,address models.AddAddress)error{
	rslt :=usrU.userRepo.CheckIfFirstAddress(id)
	var checkAddress bool 
	if !rslt {
		checkAddress =true

	}else{
		checkAddress=false
	}
	if err :=usrU.userRepo.AddAddress(id,address,checkAddress);err !=nil{
		return err
	}
	return nil
}
func (usrU *userUsecase)GetAddresses(id int )([]domain.Address,error){
	addresses,err := usrU.userRepo.GetAddresses(id)
	if err !=nil{
		return []domain.Address{},err
	}
	return addresses,nil
}

func( usrU *userUsecase)GetUserDetails(id int)(models.UserResponse,error){
	userDetails,err :=usrU.userRepo.GetUserDetails(id)
	if err !=nil{
		return models.UserResponse{},err
	}
	return userDetails,nil
}
func (usrU *userUsecase)ChangePassword(id int ,old string,password string,repassword string)error{
	userPass,err :=usrU.userRepo.GetPassword(id)
	if err !=nil{
		return errors.New("couldn't get user password")
	}
	if err :=bcrypt.CompareHashAndPassword([]byte(userPass),[]byte(old));err !=nil{
		return errors.New("password is incorrect ")

	}
	if password !=repassword{
		return errors.New("password not matching")
	}
	newPass ,err :=bcrypt.GenerateFromPassword([]byte (password),10)
	if err !=nil{
		return err
	}
	return usrU.ChangePassword(id ,string(newPass))
}
func(usrU *userUsecase) GetCartID(userID int)(int,error){
	cartId,err :=usrU.userRepo.GetCartID(userID)
	if err !=nil{
		return 0,errors.New("couldn't get cart id")
	}
	return cartId,nil
}
func(usrU *userUsecase) EditUser(id int ,userData models.EditUser)error{

	if userData.Name!= "" && userData.Name !="string"{
		err :=usrU.userRepo.EditName(id,userData.Name)
		if err !=nil{
			return err
		}
	}
	if userData.Email!= "" &&userData.Email!="string"{
		err :=usrU.userRepo.EditEmail(id,userData.Email)
		if err != nil{
			return err
		}
	}
	if userData.Phone != "" && userData.Phone != "string"{
		err := usrU.userRepo.EditPhone(id,userData.Phone)
		if err !=nil{
			return err
		}
	}
	if userData.Username!= "" && userData.Username!= "string"{
		err := usrU.userRepo.EditUsername(id,userData.Username)
		if err !=nil{
			return err
		}
	}
	return nil
}
func (usrU *userUsecase)GetCart(id ,page,limit int) ([]models.GetCart,error)  {
	//find cart id
	cartId,err :=usrU.GetCartID(id)
	if err !=nil{
		return []models.GetCart{},errors.New("could't find cart id")
	}
	//find products inside cart
	products,err :=usrU.userRepo.GetProductsInCart(cartId,page,limit)
	if err !=nil{
		return []models.GetCart{},errors.New("couldn't find products in cart")
	}
	//find products name

	var productsName []string
	for i :=range products{
		prdName,err :=usrU.userRepo.FindProductNames(products[i])

		if err !=nil{
			return []models.GetCart{},err
		}
		productsName =append(productsName, prdName)
	}
	//find quantity
	

	
}