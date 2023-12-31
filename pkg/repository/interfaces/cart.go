package interfaces

import "ecommerce/pkg/domain"

type CartRepository interface {
	GetAddresses(id int) ([]domain.Address, error)
	CheckIfInvAdded(invId,cartId int ) bool
	GetCartId(user_id int )(int, error)
	CreateNewCart(user_id int )(int, error)
	AddLineItems(invId, cartId int ) error
	AddQuantity(invId, cartId int ) error
}
