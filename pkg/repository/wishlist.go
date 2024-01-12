package repository

import (
	"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type wishlistRepository struct{
	DB *gorm.DB
}
//constructor function

func NewWishlistRepository(DB *gorm.DB) interfaces.WishListRepository{
	return &wishlistRepository{
		DB: DB,
	}
}
func (wlr *wishlistRepository) GetWishListId(user_id int)(int ,error){
	var wishlistId int
	if err :=wlr.DB.Raw("SELECT id FROM wishlist WHERE user_id=?",user_id).Scan(&wishlistId).Error;err!=nil{
		return 0,errors.New("wishlist id not found")

	}
	return wishlistId,nil
}
func (wlr *walletRepository)GetWishlist(id int) ([]models.GetWishlist,error){
	var getWishlist []models.GetWishlist
  
	query :=`
	 
	SELECT wishlist.user_id ,categories.category,inventories.product_name,inventories.price
	FROM wishlist
	JOIN wishlist_items.wishlist_id=wishlist.idgetWishlist
	JOIN inventories ON wishlist_items.inventory_id=inventories.id
	JOIN categories ON inventories.category_id=categories.id
	WHERE wishlist.user_id
	
	`
	if err :=wlr.DB.Raw(query,id).Scan(&getWishlist).Error;err !=nil{
		return []models.GetWishlist{},err
	}
	return getWishlist,nil
}

