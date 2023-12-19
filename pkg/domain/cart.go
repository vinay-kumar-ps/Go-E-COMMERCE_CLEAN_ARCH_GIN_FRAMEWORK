package domain

//cart  represents the cart of the user

type Cart struct {
	ID     uint `json:"id" gorm:"primarykey"`
	UserId uint `json:"user_id" gorm:"not null"`
	User   User `json:"-" gorm:"foreignkey:UserID"`
}

//LineItems represents products in the cart of user
type LineItems struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	CartID      uint      `json:"cart_id" gorm:"not null"`
	Cart        Cart      `json:"-" gorm:"foriegnkey:CartId "`
	InventoryID uint      `json:"inventory_id" gorm:"not null"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int       `json:"quantity" gorm:"default:1"`
}
