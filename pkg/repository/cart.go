package repository

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &cartRepository{
		DB: DB,
	}

}
func (cr *cartRepository) GetAddresses(id int) ([]domain.Address, error) {
	var Addresses []domain.Address
	err := cr.DB.Raw("SELECT * FROM ADDRESSES WHERE id =?", id).Scan(&Addresses).Error
	if err != nil {
		return []domain.Address{}, err

	}
	return Addresses, nil
}
func (cr *cartRepository) GetCartId(user_id int) (int, error) {
	var userId int
	err := cr.DB.Raw("SELECT ID FROM CARTS WHERE user_id =?", user_id).Scan(&userId).Error
	if err != nil {
		return 0, err
	}
	return userId, nil
}
func (cr *cartRepository) CreateNewCart(user_id int) (int, error) {
	var id int
	err := cr.DB.Exec(`INSERT INTO carts (user_id) VALUES (?)`, user_id).Error
	if err != nil {
		return 0, err
	}
	return id, nil

}
func (cr *cartRepository) AddLineItems(inventory_id, cart_id int) error {
	err := cr.DB.Exec(`INSERT INTO line_items (inventory_id,cart_id) VALUES (?,?)`, inventory_id, cart_id).Error
	if err != nil {
		return err
	}
	return nil
}
func (cr *cartRepository) CheckIfInvAdded(InvId, cartId int) bool {
	var count int = 0
	err := cr.DB.Raw("SELECT COUNT(id) FROM line_items where cart_id=? and inventory_id =?", cartId, InvId).Scan(&count).Error
	if err != nil {
		return false
	}
	if count < 1 {
		return false
	}
	return true
}
func (cr *cartRepository) AddQuantity(InvId, cartId int) error {
	err := cr.DB.Raw("update line_items set quantity=quantity+1 where cart_id = ? and inventory_id = ?", cartId, InvId).Error
	if err != nil {
		return err

	}
	return nil
}
