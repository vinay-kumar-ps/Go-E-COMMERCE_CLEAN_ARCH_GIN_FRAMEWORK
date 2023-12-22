package models

type GetWishlist struct {
	ProductName string  `json:"Product_name"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}
