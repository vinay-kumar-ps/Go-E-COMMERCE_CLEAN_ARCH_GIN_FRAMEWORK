package domain

//Inventory represents the products in the website

type Inventory struct {
}

// category of product
type Category struct {
	ID       int    `json:"id" gorm:"unique;not null"`
	Category string `json:"category"`
	// Image Image `json:"category_image"`
}

type Image struct {
	ID          int       `json:"id" gorm:"primarykey;not null"`
	InventoryID int       `json:"inventory_id"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	ImageUrl    string    `json:"image_url"`
}