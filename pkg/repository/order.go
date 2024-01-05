package repository

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/repository/interfaces"
	"ecommerce/pkg/utils/models"
	"errors"
	"time"

	"gorm.io/gorm"
)


type orderRepository struct{
	DB *gorm.DB
}

//constructor function
func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{
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
func (orr *orderRepository) GetProductNameFromId(id int) (string,error) {
	var productName  string
	err := orr.DB.Raw("SELECT product_name FROM inventories WHERE id =?", id).Scan(&productName).Error
	if err!=nil{
		return "",err
	}
	return productName,nil
}

func(orr *orderRepository) OrderItems(userid int,order models.Order,total float64) (int,error) {

	var id int 

	query :=`
	 
	    INSERT INTO orders
		(user_id,address_id,price,payment_method_id,ordered_at)

		VALUES 
		(?,?,?,?,?)
		RETURNING id
	`
	err:= orr.DB.Raw(query,userid,order.AdressId,total,order.PaymentID,time.Now()).Scan(&id).Error

	if err!=nil{
		return 0,nil
	}
	return id,nil
}
func (orr *orderRepository)  AddOrderProducts(order_id int ,cart[]models.Getcart)error{
	query :=`
	 
	  INSERT INTO 
	         order_items
			 (order_id,inventory_id,quantity,total_price)
			 VALUES 
                 (?,?,?,?)

	`
         for _, cartvals :=range cart {
			var invId int
			err := orr.DB.Raw("SELECT id FROM inventories WHERE product_name=?",cartvals.ProductName).Scan(&invId).Error

			if err != nil{
			return err
			}
			if err := orr.DB.Raw(query ,order_id,invId,cartvals.Quantity,cartvals.Total).Error;err !=nil {
				return err

			}
		 }
		 return nil

}

func(orr *orderRepository) CancelOrder (orderid int) error{
	err :=orr.DB.Exec("UPDATE orders SET order_status='CANCELED' WHERE id=?",orderid ).Error

	if err!=nil{
		return err
	}
	return nil
}
func (orr *orderRepository) EditOrderStatus (status string, id int )error{
	err := orr.DB.Exec("UPDATE orders SET order_status =? WHERE id =?",status,id).Error
	if err!= nil {
		return err
	}
	return nil
}
func (orr *orderRepository) MarkAsPaid(orderID  int )error{
	if err :=orr.DB.Exec("UPDATE orders SET payment_status ='PAID' WHERE id=?",orderID).Error;err!=nil{
		return err
	}
	return nil
}
func (orr *orderRepository) AdminOrders (page ,limit int ,status string)([]domain.OrderDetails,error) {

	if page == 0{
		page = 1
	}
	if limit == 0{
		limit=10
	}
	offset := (page -1 )* limit

	var orderDetails []domain.OrderDetails

	query :=`
	SELECT
	    orders.id AS order_id,users.name AS username,
		CONTACT
		    (addresses.house_name, '',adresses.street,'',adresses.city,'')
			AS address,payment_methods AS payment_method,orders .price As total
			FROM 
			    orders

				JOIn
				   users

				   ON
				    users.id=orders.id

					JOIN
					  adresses

					  ON
					   orders.adress_id=addresses.id

					   JOIN
					     payment_methods 

						 ON
						  orders.payment_method_id=payment_methods_id
						   
						  WHERE
						     order_status=? limit ? offset ?

							
	`
	err :=orr.DB.Raw(query,status,limit,offset).Scan(&orderDetails).Error
	if err != nil{
		return[]domain.OrderDetails{},err
	}
	return orderDetails,nil

}
func(orr *orderRepository) CheckOrder (orderID string,userID int)  error{

	var Count  int
	err :=orr.DB.Raw("SELECT COUNT (*)FROM orders WHERE order_id=?",orderID).Scan(&Count).Error
	if err!=nil{
		return err
	}
	if Count < 0 {
		return  errors.New("no such orders exixt")
	}
	var checkUser int
		chUser := orr.DB.Raw("SELECT user_id FROM orders WHERE order_id=?",orderID).Scan(&checkUser)
		if err!=nil{

			return chUser.Error
		}
		if userID!=checkUser{

			return errors.New("no order did by this user")

		}
		return nil

} 
func (orr *orderRepository) GetOrderDetails (orderiD string) (domain.Order,error) {


	 var  orderDetails domain.Order

	 err := orr.DB.Raw("SELECT * FROM orders WHERE order_id=?",orderiD).Scan(&orderDetails).Error

	 if  err != nil{
		return domain.Order{},err
	 }
	 return orderDetails,nil

}

func (orr *orderRepository) FindAmountOrderID (orderID int ) (float64,error){
	var amount float64
	err := orr.DB.Raw("SELECT price FROM orders WHERE orders_id=?",orderID).Scan(&amount).Error
	 if err !=nil{
		return 0,err
	 }
	 return amount,nil
}

func (orr *orderRepository) FindUserIdFromOrderID(orderID int) (int,error){
	 var userId int 

	 err := orr.DB.Raw("SLECT user_id FROM orders HWERE order_id =?",orderID).Scan(&userId).Error
	  if err != nil{

		return 0,err
	  }
	return userId,nil
}

func(orr *orderRepository) ReturnOrder(id int )error{
	 err :=orr.DB.Exec("UPDATE orders SET  order_status='RETURNED' WHERE id=?",id).Error
	
	 if err!=nil{
   return err
	 }
	 return nil
}

func (orr *orderRepository) CheckOrderStatus(orderID int ) (string ,error){
var orderStatus string

err := orr.DB.Raw("SELECT order_status FROM orders WHERE order_id =?",orderID).Scan(&orderStatus).Error


if err != nil{
 return "",err
}
return orderStatus,nil
}

func(orr *orderRepository) CheckPaymentStatus (orderID int) (string ,error){
	var paymentStatus string
	err := orr.DB.Raw("SELECT payment_status FROM orders WHERE order_id=?",orderID).Scan(&paymentStatus).Error
	if err != nil{
		return "", err
	}
	return paymentStatus ,nil
}