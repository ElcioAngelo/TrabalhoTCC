package model

import "time"


type Order struct {
	ID             int       `json:"id"`               	// * Unique identifier for the order
	OrderDate      time.Time `json:"order_date"`        // * Date the order was placed
	ProductQuantity int      `json:"product_quantity"`  // * Quantity of the product ordered
	Status         string    `json:"status"`            // * Status of the order (e.g., "pending", "shipped")
	PaymentMethod  string    `json:"payment_method"`    // * Method used to pay for the order (e.g., "credit card", "paypal")
	ItemOrderID    int       `json:"item_order_id"`     // * ID of the item within the order (if multiple items)
}