package model

type User struct {
	ID              int       `json:"id"`                // * Unique identifier for the user
	Name            string    `json:"name"`              // * Name of the user
	Email           string    `json:"email"`             // * Email of the user
	Password        string    `json:"password"`          // * User's password (in a real application, ensure it's hashed)
	CellphoneNumber string    `json:"cellphone_number"`  // * Cellphone number of the user
	ShippingAddress string    `json:"shipping_address"`  // * Shipping address for delivery
	PaymentAddress  string    `json:"payment_address"`   // * Payment address (e.g., billing address)       // * Date when the user was created
}

