package model

import "time"

type Order struct {
	ID            int       `json:"id"`
	OrderDate     time.Time `json:"order_date"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	Username      string    `json:"username"`
}
