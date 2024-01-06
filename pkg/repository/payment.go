package repository

import (
	"ecommerce/pkg/domain"
	"ecommerce/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type PaymentRepository struct{
	DB *gorm.DB
}

//constructor function

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentRepository{
		DB: DB,
	}
}

func (pr *PaymentRepository) AddNewPaymentMethod (paymentMethod string)error{
	query := `INSERT INTO payment_method (payment_method)VALUES(?)`
	if err:= pr.DB.Exec(query ,paymentMethod).Error; err != nil{
		
		return err 
	}
	return nil
}

func (pr *PaymentRepository) RemovePaymentMethod(paymentMethodId int ) error{
	query := `DELETE FROM payment_methods WHERE id =?`

	if err:= pr.DB.Exec(query ,paymentMethodId).Error; err != nil{
		
		return err 
	}
	return nil

}

func (pr *PaymentRepository)GetPaymentMethods()([]domain.PaymentMethod,error){
	var paymentMethods []domain.PaymentMethod

	err :=pr.DB.Raw("SELECT *FROM payment_method").Scan(&paymentMethods).Error
	if err!=nil{
		return []domain.PaymentMethod{},err
	}
	return paymentMethods,nil
}
func (pr *PaymentRepository) FindUsername(user_id int ) (string,error){
	var userName string
	err := pr.DB.Raw("SELECT name FROM users WHERE id=?",user_id).Scan(&userName).Error
	if err !=nil{
		return "",err
	}
	return userName,nil
}
func (pr *PaymentRepository) FindPrice (order_id int ) (float64,error){
	var price float64

	err := pr.DB.Raw("SELECT price FROM orders WHERE id=?",order_id).Scan(&price).Error
	if err !=nil{
		return 0,err
	}
	return price,nil
}

func (pr *PaymentRepository) updatePaymentDetails (orderId ,paymentId ,razorId string)error{
	status :="paid"
	 query := `UPDATE orders SET payment_status =$1,payment_id=$3 WHERE id $2`
	 if err := pr.DB.Exec(query ,status,orderId,paymentId).Error;err !=nil{
		return err
	 }
	 return nil
}