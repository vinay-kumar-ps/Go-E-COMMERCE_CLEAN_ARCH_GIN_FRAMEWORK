package domain

//cart represents the wishlist of the user

type Wishlist struct {
	ID     uint `json:"id" gorm:"primarykey"`
	UserID uint `json:"user_id" gorm:"not null"`
	User   User `json:"-" gorm:"foreignkey:UserID"`
}

//LIneitems represents products in the wishlist of user

type WishlistItems struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	WishlistID  uint      `json:"card_id" gorm:"not null"`
	Wishlist    Wishlist  `json:"-" gorm:"foreignkey:WishListID"`
	InventoryID uint      `json:"inventory_id" gorm:"not null"`
	Inventory   Inventory `json:"-" gorm:"foriegnkey:InventoryID"`
}
