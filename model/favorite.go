package model

type Favorite struct {
	ID int 				`json:"id"`
	ProductID int		`json:"product_id"`
	UserID 	int			`json:"user_id"`
}
