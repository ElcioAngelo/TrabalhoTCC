package model 

type Product struct {
	ID int 			`json:"id_product"`
	Name string 	`json:"name"`
	Price float64	`json:"price"`
	Description string `json:"description"`
	CategoryID int		`json:"category_id"`
	BrandID int			`json:"brand_id"`
}
