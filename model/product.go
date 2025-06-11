package model 

type Product struct {
	ID 				int 		`json:"id_product"`
	Name 			string 		`json:"name"`
	Price 			float64		`json:"price"`
	Description 	string 		`json:"description"`
	Category 		string		`json:"product_category"`
	Brand			string		`json:"product_brand"`
	Product_Status 	string		`json:"product_status"`
}
