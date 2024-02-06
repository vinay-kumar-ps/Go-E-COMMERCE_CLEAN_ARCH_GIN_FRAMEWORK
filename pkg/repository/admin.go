package repository

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ad *adminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := ad.DB.Raw("select * from admins where username = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}

	return adminCompareDetails, nil
}

func (ad *adminRepository) GetUserByID(id string) (domain.Users, error) {

	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.Users{}, err
	}

	var count int
	if err := ad.DB.Raw("select count(*) from users where id = ?", user_id).Scan(&count).Error; err != nil {
		return domain.Users{}, err
	}
	if count < 1 {
		return domain.Users{}, errors.New("user for the given id does not exist")
	}

	query := fmt.Sprintf("select * from users where id = '%d'", user_id)
	var userDetails domain.Users

	if err := ad.DB.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.Users{}, err
	}

	return userDetails, nil
}

// function which will both block and unblock a user
func (ad *adminRepository) UpdateBlockUserByID(user domain.Users) error {

	err := ad.DB.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error
	if err != nil {
		return err
	}

	return nil

}

func (ad *adminRepository) GetUsers(page int) ([]models.UserDetailsAtAdmin, error) {
	// pagination purpose -
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * 5
	var userDetails []models.UserDetailsAtAdmin

	if err := ad.DB.Raw("select id,name,email,phone,blocked from users limit ? offset ?", 20, offset).Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

func (i *adminRepository) NewPaymentMethod(pay string) error {

	if err := i.DB.Exec("insert into payment_methods(payment_name)values($1)", pay).Error; err != nil {
		return err
	}

	return nil

}

func (a *adminRepository) ListPaymentMethods() ([]domain.PaymentMethod, error) {
	var model []domain.PaymentMethod
	err := a.DB.Raw("SELECT * FROM payment_methods where is_deleted = false").Scan(&model).Error
	if err != nil {
		return []domain.PaymentMethod{}, err
	}

	return model, nil
}

func (a *adminRepository) CheckIfPaymentMethodAlreadyExists(payment string) (bool, error) {
	var count int64
	err := a.DB.Raw("SELECT COUNT(*) FROM payment_methods WHERE payment_name = $1 and is_deleted = false", payment).Scan(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (a *adminRepository) DeletePaymentMethod(id int) error {
	err := a.DB.Exec("UPDATE payment_methods SET is_deleted = true WHERE id = $1 ", id).Error
	if err != nil {
		return err
	}

	return nil
}