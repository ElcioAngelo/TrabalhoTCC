package model

type UserInformation struct {
	ID int 							`json:"id"`
	UserID int						`json:"user_id"`
	Email string					`json:"email"`
	Password string					`json:"password"`
	Cellphone_number string			`json:"cellphone_number"`
	ShippingAddress string 			`json:"shipping_addres"`
	PaymentAddress string			`json:"payment_address"`
}
