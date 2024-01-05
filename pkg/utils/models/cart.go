package models

import "ecommerce/pkg/domain"

type GetCart struct {
	ProductName string `json:"product_name"`
	Category_id int `json:"category_id"`
	Quantity int `json:"quantity"`
	Total float64      `json:"total_price"`
	 DiscountedPrice float64 `json:"discounted_price"`
}
type CheckOut struct{
	Adressess []domain.Address
	Products []GetCart
	PaymentMethod[]domain.PaymentMethod
	TotalPrice float64
}
type Order struct{
	AddressId int `json:"address_id"`
	PaymentId int `json:"payment_ID"`

}
