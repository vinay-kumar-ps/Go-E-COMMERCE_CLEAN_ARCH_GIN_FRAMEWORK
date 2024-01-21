package pdf

type InvoiceData struct{
	Title  string
	Quantity int 
	price  int
   TotalAmount int
}
func (d *InvoiceData) CalculateTotalAmount ()int {
	totalAmount :=d.Quantity *d.price
	return totalAmount
}

func (d *InvoiceData) ReturnItemTotalAmount()float64{
	totalAmount :=d.CalculateTotalAmount()
	converted :=float64(totalAmount)/100
	return converted
}
func (d *InvoiceData)ReturnItemPrice()float64{
	returnPrice :=float64(d.price)/100
	return returnPrice
}
