package models

type CreateOffer struct{
	CategoryID int `json:"category_id"`
	Discount   int `json:"discount"`
}
