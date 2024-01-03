package repository

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
	"time"

	"gorm.io/gorm"
)


type orderRepository struct{
	DB *gorm.DB
}

//constructor function
func NewOrderRepository (DB *gorm.DB) interfaces.OfferRepository {
	return &orderRepository {
		DB: DB,
	}
}


func (orr *orderRepository) GetOrders(id ,page ,limit int) ([]domain.Order,error){

	if page ==0{
		page =1
	}
	if limit == 0{
		limit=10
	}

	offset :=(page -1) *limit


	var getOrders []domain.Order
	 
	err :=orr.DB.Raw("SELECT  *FROM orders WHERE id=? limit ?",id ,limit,offset).Scan(&getOrders).Error

	if err !=nil{
		return []domain.Order{},err
	}
	return getOrders,nil

}

func (orr *orderRepository) GetOrdersInRange(startDate, endDate time.Time) ([]domain.Order,error) {

	var getOrdersInTimeRange []domain.Order

	//to fetch orders with in a time range
	err := orr.DB.Raw("SELECt * FROM  orders WHERE orderd_at BETWEEN ? AND ?",startDate,endDate).Scan(&getOrdersInTimeRange).Error

	if err!=nil{
		return []domain.Order{},err
	}
	return getOrdersInTimeRange,nil
}

func (orr *orderRepository) GetProductsQuantity()([]domain.ProductReport,error){

   var GetProductQuantity	[]domain.ProductReport

   err :=orr.DB.Raw("SELECT inventory_id,quantity FROM order_items").Scan(&GetProductQuantity).Error
    if err!=nil{
		return []domain.ProductReport{},err
	}
	return GetProductQuantity,nil
		
		
	}	
	func (orr *orderRepository) GetCart(userid int) (models.Getcart,error)  {
		var cart models.Getcart
		err :=orr.DB.Raw("SELECT inventories.product_name,cart_products.quantity.cart_products.total_price AS total FROM cart_products JOIN inventories ON cart_products.inventories.id WHERE user_id=?",userid).Scan(&cart).Error

		if err != nil{

			return models.Getcart{}, err
		}
		return cart, nil

}
