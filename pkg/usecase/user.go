package usecase

import (
	config "ecommerce/pkg/config"
	"ecommerce/pkg/domain"
	helper_interface "ecommerce/pkg/helper/interface"
	interfaces "ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"
)

type userUseCase struct {
	userRepo            interfaces.UserRepository
	cfg                 config.Config
	otpRepository       interfaces.OtpRepository
	inventoryRepository interfaces.InventoryRepository
	orderRepository     interfaces.OrderRepository
	helper              helper_interface.Helper
}

func NewUserUseCase(repo interfaces.UserRepository, cfg config.Config, otp interfaces.OtpRepository, inv interfaces.InventoryRepository, order interfaces.OrderRepository, h helper_interface.Helper) *userUseCase {
	return &userUseCase{
		userRepo:            repo,
		cfg:                 cfg,
		otpRepository:       otp,
		inventoryRepository: inv,
		orderRepository:     order,
		helper:              h,
	}
}

var InternalError = "Internal Server Error"
var ErrorHashingPassword = "Error In Hashing Password"

func (u *userUseCase) UserSignUp(user models.UserDetails, ref string) (models.TokenUsers, error) {
	// Check whether the user already exist. If yes, show the error message, since this is signUp
	userExist := u.userRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")
	}
	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	referenceUser, err := u.userRepo.FindUserFromReference(ref)
	if err != nil {
		return models.TokenUsers{}, errors.New("cannot find reference user")
	}

	// Hash password since details are validated

	hashedPassword, err := u.helper.PasswordHashing(user.Password)
	if err != nil {
		return models.TokenUsers{}, errors.New(ErrorHashingPassword)
	}

	user.Password = hashedPassword

	referral, err := u.helper.GenerateRefferalCode()
	if err != nil {
		return models.TokenUsers{}, errors.New(InternalError)
	}

	// add user details to the database
	userData, err := u.userRepo.UserSignUp(user, referral)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not add the user")
	}

	// crete a JWT token string for the user
	tokenString, err := u.helper.GenerateTokenClients(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	//credit 20 rupees to the user which is the source of the reference code
	if err := u.userRepo.CreditReferencePointsToWallet(referenceUser); err != nil {
		return models.TokenUsers{}, errors.New("error in crediting gift")
	}

	//create new wallet for user
	if _, err := u.orderRepository.CreateNewWallet(userData.Id); err != nil {
		return models.TokenUsers{}, errors.New("errror in creating new wallet")
	}
	return models.TokenUsers{
		Users: userData,
		Token: tokenString,
	}, nil
}

func (u *userUseCase) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {

	// checking if a username exist with this email address
	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")
	}

	isBlocked, err := u.userRepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.TokenUsers{}, errors.New(InternalError)
	}

	if isBlocked {
		return models.TokenUsers{}, errors.New("user is blocked by admin")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, errors.New(InternalError)
	}

	err = u.helper.CompareHashAndPassword(user_details.Password, user.Password)
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse

	userDetails.Id = int(user_details.Id)
	userDetails.Name = user_details.Name
	userDetails.Email = user_details.Email
	userDetails.Phone = user_details.Phone

	tokenString, err := u.helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token")
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil

}

func (i *userUseCase) AddAddress(id int, address models.AddAddress) error {

	rslt := i.userRepo.CheckIfFirstAddress(id)
	var result bool

	if !rslt {
		result = true
	} else {
		result = false
	}

	err := i.userRepo.AddAddress(id, address, result)
	if err != nil {
		return errors.New("error in adding address")
	}

	return nil

}

func (i *userUseCase) GetAddresses(id int) ([]domain.Address, error) {

	addresses, err := i.userRepo.GetAddresses(id)
	if err != nil {
		return []domain.Address{}, errors.New("error in getting addresses")
	}

	return addresses, nil

}

func (i *userUseCase) GetUserDetails(id int) (models.UserDetailsResponse, error) {

	details, err := i.userRepo.GetUserDetails(id)
	if err != nil {
		return models.UserDetailsResponse{}, errors.New("error in getting details")
	}

	return details, nil

}

func (i *userUseCase) ChangePassword(id int, old string, password string, repassword string) error {

	userPassword, err := i.userRepo.GetPassword(id)
	if err != nil {
		return errors.New(InternalError)
	}

	err = i.helper.CompareHashAndPassword(userPassword, old)
	if err != nil {
		return errors.New("password incorrect")
	}

	if password != repassword {
		return errors.New("passwords does not match")
	}

	newpassword, err := i.helper.PasswordHashing(password)
	if err != nil {
		return errors.New("error in hashing password")
	}

	return i.userRepo.ChangePassword(id, string(newpassword))

}

func (u *userUseCase) ForgotPasswordSend(phone string) error {

	ok := u.otpRepository.FindUserByMobileNumber(phone)
	if !ok {
		return errors.New("the user does not exist")
	}

	u.helper.TwilioSetup(u.cfg.ACCOUNTSID, u.cfg.AUTHTOKEN)
	_, err := u.helper.TwilioSendOTP(phone, u.cfg.SERVICESID)
	if err != nil {
		return errors.New("error ocurred while generating OTP")
	}

	return nil

}

func (u *userUseCase) ForgotPasswordVerifyAndChange(model models.ForgotVerify) error {
	u.helper.TwilioSetup(u.cfg.ACCOUNTSID, u.cfg.AUTHTOKEN)
	err := u.helper.TwilioVerifyOTP(u.cfg.SERVICESID, model.Otp, model.Phone)
	if err != nil {
		return errors.New("error while verifying")
	}

	id, err := u.userRepo.FindIdFromPhone(model.Phone)
	if err != nil {
		return errors.New("cannot find user from mobile number")
	}

	newpassword, err := u.helper.PasswordHashing(model.NewPassword)
	if err != nil {
		return errors.New("error in hashing password")
	}

	// if user is authenticated then change the password i the database
	if err := u.userRepo.ChangePassword(id, string(newpassword)); err != nil {
		return errors.New("could not change password")
	}

	return nil
}

func (i *userUseCase) EditName(id int, name string) error {

	err := i.userRepo.EditName(id, name)
	if err != nil {
		return errors.New("could not change")
	}

	return nil

}

func (i *userUseCase) EditEmail(id int, email string) error {

	err := i.userRepo.EditEmail(id, email)
	if err != nil {
		return errors.New("could not change")
	}

	return nil

}

func (i *userUseCase) EditPhone(id int, phone string) error {

	err := i.userRepo.EditPhone(id, phone)
	if err != nil {
		return errors.New("could not change")
	}

	return nil

}

func (u *userUseCase) GetCart(id int) (models.GetCartResponse, error) {

	//find cart id
	cart_id, err := u.userRepo.GetCartID(id)
	if err != nil {
		return models.GetCartResponse{}, errors.New(InternalError)
	}
	//find products inide cart
	products, err := u.userRepo.GetProductsInCart(cart_id)
	if err != nil {
		return models.GetCartResponse{}, errors.New(InternalError)
	}
	//find product names
	var product_names []string
	for i := range products {
		product_name, err := u.userRepo.FindProductNames(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		product_names = append(product_names, product_name)
	}

	//find quantity
	var quantity []int
	for i := range products {
		q, err := u.userRepo.FindCartQuantity(cart_id, products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		quantity = append(quantity, q)
	}

	var price []float64
	for i := range products {
		q, err := u.userRepo.FindPrice(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		price = append(price, q)
	}

	var images []string
	var stocks []int

	for _, v := range products {
		image, err := u.userRepo.FindProductImage(v)
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}

		stock, err := u.userRepo.FindStock(v)
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}

		images = append(images, image)
		stocks = append(stocks, stock)
	}

	var categories []int
	for i := range products {
		c, err := u.userRepo.FindCategory(products[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		categories = append(categories, c)
	}

	var getcart []models.GetCart
	for i := range product_names {
		var get models.GetCart
		get.ID = products[i]
		get.ProductName = product_names[i]
		get.Image = images[i]
		get.Category_id = categories[i]
		get.Quantity = quantity[i]
		get.Total = price[i]
		get.StockAvailable = stocks[i]
		get.DiscountedPrice = 0

		getcart = append(getcart, get)
	}

	//find offers
	var offers []int
	for i := range categories {
		c, err := u.userRepo.FindofferPercentage(categories[i])
		if err != nil {
			return models.GetCartResponse{}, errors.New(InternalError)
		}
		offers = append(offers, c)
	}

	//find discounted price
	for i := range getcart {
		getcart[i].DiscountedPrice = (getcart[i].Total) - (getcart[i].Total * float64(offers[i]) / 100)
	}

	var response models.GetCartResponse
	response.ID = cart_id
	response.Data = getcart

	//then return in appropriate format

	return response, nil

}

func (i *userUseCase) RemoveFromCart(cart, inventory int) error {

	err := i.userRepo.RemoveFromCart(cart, inventory)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) UpdateQuantityAdd(id, inv int) error {

	err := i.userRepo.UpdateQuantityAdd(id, inv)
	if err != nil {
		return err
	}

	return nil

}

func (i *userUseCase) UpdateQuantityLess(id, inv int) error {

	err := i.userRepo.UpdateQuantityLess(id, inv)
	if err != nil {
		return err
	}

	return nil

}

// func (i *userUseCase) GetMyReferenceLink(id int) (string, error) {

// 	baseURL := "jerseyhub.com/users/signup"

// 	referralCode, err := i.userRepo.GetReferralCodeFromID(id)
// 	if err != nil {
// 		return "", errors.New("error getting ref code")
// 	}

// 	referralLink := fmt.Sprintf("%s?ref=%s", baseURL, referralCode)

// 	//returning the link
// 	return referralLink, nil
// }
