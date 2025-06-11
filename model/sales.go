package model

type Sales struct {
	Name            string    `json:"name"`              // * Name of the user
	Email           string    `json:"email"`             // * Email of the user
	CellphoneNumber string    `json:"user_cellphone"`  // * Cellphone number of the user
	PaymentAddress  string    `json:"user_payment_address"`
	ShippingAddress string    `json:"user_shipping_address"`  // * Shipping address for delivery
	TotalRevenue	float32	  `json:"total_revenue"`
	SaleDate		string	  `json:"sale_date"`
}
