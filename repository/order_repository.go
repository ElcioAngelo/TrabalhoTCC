package repository

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type OrderRepository struct {
	connection *sql.DB
}


func NewOrderRepository(connection *sql.DB) OrderRepository {
	return OrderRepository{
		connection: connection,
	}
}

func (or *OrderRepository) SetUserOrder(user_id int, products []int) error {
	var OrderID int

	query := `
		INSERT INTO orders (order_date, payment_method, status, user_id)
		VALUES (NOW(), NULL, 'pending', $1)
		RETURNING id;
	`
	err := or.connection.QueryRow(query, user_id).Scan(&OrderID)
	if err != nil {
		panic(err.Error())
	}

	productCount := make(map[int]int)
	for _, productID := range products {
		productCount[productID]++
	}

	var productIDs []int
	var quantities []int
	for productID, quantity := range productCount {
		productIDs = append(productIDs, productID)
		quantities = append(quantities, quantity)
	}
	

	fmt.Println("Order ID:", OrderID)
	fmt.Println("Product IDs:", productIDs)
	fmt.Println("Quantities:", quantities)


	// Call bulk function
	_, err = or.connection.Exec(`SELECT add_products_bulk($1, $2, $3)`,
		OrderID, pq.Array(productIDs), pq.Array(quantities))
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func (or *OrderRepository) ReturnOrder(user_id int) ([]map[string]interface{}, error) {
	query := `
		SELECT o.id, o.order_date, o.status, o.user_id, p."name", op.quantity, p.price, p.price * op.quantity as "preco_total"
		FROM orders o
		INNER JOIN order_products op ON op.order_id = o.id
		INNER JOIN products p ON p.id = op.product_id
		WHERE o.user_id = $1;
	`

	// Use Query instead of Exec because we want to fetch rows
	rows, err := or.connection.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to store the results
	var results []map[string]interface{}

	// Iterate over the rows
	for rows.Next() {
		var orderID int
		var orderDate string
		var status string
		var userID int
		var productName string
		var quantity int
		var price float64
		var precoTotal float64

		// Scan each row into variables
		if err := rows.Scan(&orderID, &orderDate, &status, &userID, &productName, &quantity, &price, &precoTotal); err != nil {
			return nil, err
		}

		// Store the row in a map
		order := map[string]interface{}{
			"order_id":    orderID,
			"order_date":  orderDate,
			"status":      status,
			"user_id":     userID,
			"product_name": productName,
			"quantity":    quantity,
			"price":       price,
			"preco_total": precoTotal,
		}

		// Append the map to the results slice
		results = append(results, order)
	}

	// Check for any error after looping through rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return the results
	return results, nil
}

func (or *OrderRepository) ReturnAllOrders() ([]map[string]interface{}, error) {
	query := 
	`
		SELECT u."name", COUNT(*), SUM(p.price * op.quantity) AS total_price, o.status
FROM orders o
inner join order_products op on o.id = op.order_id
inner join products p on p.id = op.product_id
inner join users u on u.id = o.user_id
GROUP BY user_id, o.status, u."name"
ORDER BY user_id, o.status, u."name"
	`

	// Use Query instead of Exec because we want to fetch rows
	rows, err := or.connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create a slice to store the results
	var results []map[string]interface{}

	// Iterate over the rows
	for rows.Next() {
		var user_name string
		var quantity int
		var precoTotal float64
		var status string

		// Scan each row into variables
		if err := rows.Scan(&user_name, &quantity, &precoTotal, &status); err != nil {
			return nil, err
		}

		// Store the row in a map
		order := map[string]interface{}{
			"user_name":     user_name,
			"quantity":    quantity,
			"preco_total": precoTotal,
			"status": status,
		}

		// Append the map to the results slice
		results = append(results, order)
	}

	// Check for any error after looping through rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return the results
	return results, nil
}


// func (or *OrderRepository) AlterOrder(status string, order_id int) (error) {
// 	query := 
// 	`
// 		update orders
// 		set status = $1
// 		where id = $2;
// 	`
// 	_, err := or.connection.Exec(query,status,order_id);
// 	if err != nil {
// 		panic(err.Error());
// 	}

// 	return err;
// }


