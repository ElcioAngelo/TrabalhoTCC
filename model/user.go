package model

type User struct {
	ID              int       `json:"id"`                
	Name            string    `json:"name"`              
	Email           string    `json:"email"`             
	Password        string    `json:"password"`
	CellphoneNumber string 	  `json:"cellphone_number"`
	State			string 	  `json:"state"`
	PostalCode		string 	  `json:"postal_code"`
	City 			string 	  `json:"city"`
	Address			string    `json:"address"`
	AddressNumber	string	  `json:"address_number"`
	UserRole		string	  `json:"user_role"`
}
 
// type User struct {
//     UserID         int    `json:"user_id"`
//     Name           string `json:"name"`
//     Email          string `json:"email"`
//     Password       string `json:"password"`
//     CellphoneNumber string `json:"cellphone_number"`
//     Address        string `json:"address"`
//     AddressNumber  string `json:"address_number"`
//     City           string `json:"city"`
//     State          string `json:"state"`
//     PostalCode     string `json:"postal_code"`
// }

