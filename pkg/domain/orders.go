package domain

import "time"

//order represents the users order

type Order struct {
	ID              int           `json:"id" gorm:"primarykey;autoincrement"`
	UserID          int           `json:"user_id" gorm:"not null"`
	User            User          `json:"-" gorm:"foreignkey:UserID"`
	AddressID       int           `json:"address_id" gorm:"not null"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID int           `json:"paymentmethodID" gorm:"default:1"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foriegnkey:PaymentMethodID"`
	PaymentID       string        `json:"paymentID"`
	Price           float64       `json:"price"`
	OrderedAt       time.Time     `json:"orderedAt"`
	OrderStatus     string        `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING','SHIPPED,'DELIVERED','CANCELED','RETURNED')"`
	PaymentStatus   string        `json:"paymentStatus" gorm:"default:'pending'"`
}

//orderItem represents the product details of the order

type OrderItem struct {
	ID          int       `json:"id" gorm:"primarykey;autoIncrement"`
	OrderID     int       `json:"order_id"`
	Order       Order     `json:"-" gorm:"foreignkey:OrderID"`
	InventoryID int       `json:"inventory_id"`
	Inventory   Inventory `json:"_" gorm:"foreignkey:InventoryID"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
}

//admminOrderResponse represents the order details with order status
type AdminOrderResponse struct {
	Pending   []OrderDetails
	Shipped   []OrderDetails
	Delivered []OrderDetails
	Canceled  []OrderDetails
	ReturnedS []OrderDetails
}
//OrderDetails represents the details of order
type OrderDetails struct{
	ID int `json:"order_id" gorm:"coloumn:order_id"`
	Username string `json:"name"`
	Address string `json:"address"`
	PaymentMethod string`json:"paymentmethod"`
	//PaymentMethod PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	Total float64 `json:"total"`
}

type PaymentMethod struct {
	ID            int    `gorm:"primarykey"`
	PaymentMethod string `json:"PaymentMethod" validate:"required" gorm:"unique"`
}
type SalesReport struct{
	Orders []Order
	TotalRevenue []float64
	TotalOrders []int
BestSellers []string
}

type productReport struct{
	InventoryID int
	Quantity int
}
