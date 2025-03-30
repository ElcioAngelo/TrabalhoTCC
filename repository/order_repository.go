package repository

import (
	"database/sql"
)


type OrderRepository struct {
	connection *sql.DB
}

func NewOrderRepository(connection *sql.DB) OrderRepository{
	return OrderRepository{
		connection: connection,
	}
}

// func(or * OrderRepository) GetSoldOrders([]model.Order, error) {
// 	query := 
// 	`
// 	select * from orders o
// 	where o.status = 'sold';
// 	`

// 	result, err := or.connection.Query(query); if err != nil {
// 		panic(err)
// 	}
	
	
// } 