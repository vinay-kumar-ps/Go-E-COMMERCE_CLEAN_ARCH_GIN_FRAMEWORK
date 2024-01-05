package models

type OrderPaymentsDetails struct {
	UserID     int     `json:"user_id"`
	Username   string  `json:"username"`
	RazorId   int     `json:"razor_id"`
	OrderID    int     `json:"order_id"`
	FinalPrice float64 `json:"final_price"`
}
type InvoiceData struct {
	Title       string
	Quantity    int
	Price       int
	TotalAmount int
}
type Invoice struct{
	Name string 
	Adress string 
	InvoiceItems []*InvoiceData
}
