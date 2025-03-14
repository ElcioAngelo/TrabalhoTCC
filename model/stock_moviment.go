package model

type StockMoviment struct {
	ID	int 		`json:"id"`
	ProductID int	`json:"product_id"`
	Quantity int	`json:"quantity"`
	BatchID int		`json:"batch_id"`
	Value float32	`json:"value"`
	Operation string `json:"operation"`
}
