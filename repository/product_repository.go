package repository

import (
	"database/sql"
	"fmt"

	"github.comElcioAngelo/TrabalhoTCC.git/model"
)

type ProductRepository struct {
	connection *sql.DB
}


func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := `SELECT * FROM products;`

	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
			&productObj.Description,
			&productObj.CategoryID,
			&productObj.BrandID,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (error) {

	query := `insert into Products values
	($1,$2,$3,$4,$5)`
	
	
	result, err := pr.connection.Exec(query,product.Name,
		product.Price,
		product.Description,
		product.CategoryID,
		product.BrandID)	
 	if err != nil {
		panic(err)
	}

	affectedRows, _ := result.RowsAffected()
	fmt.Printf("Rows affeceted: %d", affectedRows)
	return err 
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(id_product).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}

func joinStrings(arr []string, seperator string) string {
	return fmt.Sprintf("%s", arr)
}

func (pr *ProductRepository) EditProduct(fields []string, values []interface{},id string) (error) {
	query := fmt.Sprintf("update Product set %s where id = ?", joinStrings(fields, ", "))
	values = append(values, id)

	_, err := pr.connection.Exec(query, values...)
	if err != nil {
		panic(err)
	}
	return err
}

