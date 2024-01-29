package models

type InventoryResponse struct {
	ProductID int
	//stock int
}

type Inventory struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	Image       string  `json:"image"`
	ProductName string  `json:"productname"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
type UpdateInventory struct {
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"productName"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
type InventoryList struct {
	ID          uint    `json:"id"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	ProductName string  `json:"productName"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
type InventoryDetails struct{
	Inventory Inventory
	AdditionalImages []ImagesInfo
}
type ImagesInfo struct{
	ID int `json:"id"`
	Imageurl string `json:"imageurl"`
}