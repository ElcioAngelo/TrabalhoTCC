package model 

type Product struct {
	ID 				int 		`json:"id_product"`
	Name 			string 		`json:"name"`
	Price 			float64		`json:"price"`
	Description 	string 		`json:"description"`
	Category 		int			`json:"product_category"`
	Brand			int			`json:"product_brand"`
}
