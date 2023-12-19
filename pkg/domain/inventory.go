package domain
//Inventory represents the products in the website

type Inventory struct {
	ID          uint     `json:"id" gorm:"unique;not null"`
	CategoryID  int      `json:"category_id"`
	Category    Category `json:"-" gorm:"foreignkey:CategoryID;constraint:OneDelete:CASCADE"`
	ProductName string   `json:"product_name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Stock       int      `json:"stock"`
	Price       int      `json:"price"`
}

//category represents the category of product
type Category struct {
	ID uint`json:"id" gorm:"unique;not null"`
	Category string `json:"category"`
}
 type Image struct{
	ID uint `json:"id" gorm:"unique;not null"`
	InventoryID int `json:"inventory_id"`
	Inventory int 	`json:"-" gorm:"foreignkey:InventoryiD"`
	ImageURL string `json:"imageurl"`
 }