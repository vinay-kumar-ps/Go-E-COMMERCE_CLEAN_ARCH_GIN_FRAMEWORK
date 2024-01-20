package usecase

import (
	
	"errors"

  interfaces"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
 services"ecommerce/pkg/usecase/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type adminUsecase struct {
	adminRepository interfaces.AdminRepository
}

// constructor function
func NewAdminUsecase(adRepo interfaces.AdminRepository) services.AdminUsecase {
	return &adminUsecase{
		adminRepository: adRepo,
	}
}

func (au *adminUsecase) LoginHandler(adminDetails models.AdminLogin) (models.AdminToken, error) {
	// Getting details of the admin based on the email provided

	adminCompareDetails, err := au.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return models.AdminToken{}, errors.New("admin not found")
	}

	hash, err := helper.PasswordHashing(adminDetails.Password)
	if err != nil {
		return models.AdminToken{}, err
	}
	// Compare password from database that provided by admin
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(adminDetails.Password))
	if err != nil {
		return models.AdminToken{}, err
	}

	token, err := helper.GenerateAdminToken(adminCompareDetails)
	if err != nil {
		return models.AdminToken{}, err
	}
	return models.AdminToken{
		Username: adminCompareDetails.Username,
		Token:    token,
	}, nil
}
func (au *adminUsecase) BlockUser(id string) error {
	user, err := au.adminRepository.GetUserById(id)
	if err != nil {
		return errors.New("user not found")

	}
	if !user.Permission {
		return errors.New("already blocked")
	} else {
		user.Permission = false
	}
	err = au.adminRepository.UpdateBlockUserById(user)
	if err != nil {
		return err
	}
	return nil

}

func (au *adminUsecase) UnblockUser(id string) error {
	user, err := au.adminRepository.GetUserById(id)
	if err != nil {
		return errors.New("user not found")
	}
	if !user.Permission {
		// means that user is blocked(false)..then,
		user.Permission = true
	}
	err = au.adminRepository.UpdateBlockUserById(user)
	if err != nil {
		return errors.New("error in unblock user")
	}
	return nil
}
func (au *adminUsecase) Getusers(page, limit int) ([]models.UserDetailsAtAdmin, error) {
	users, err := au.adminRepository.GetUsers(page, limit)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, errors.New("getting users failed")

	}
	return users, nil
}
