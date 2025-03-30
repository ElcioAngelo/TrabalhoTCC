package repository

import (
	"database/sql"
	"fmt"

	"trabalhoTcc.com/mod/model"
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

	query := `select p."name", p.description, c."name" as "product_category", b."name" as "product_brand"
	from products p
	inner join categories c on p.category_id  = c.id
	inner join brands b on p.brand_id  = b.id;`

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
			&productObj.Category,
			&productObj.Brand,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		
		// !! Caso o status do produto seja "inativo", 
		// !! ele não é adcionado na lista,
		// !! portanto não é mostrado.
		if productObj.Product_Status != "inactive" {
			productList = append(productList, productObj)
		}
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) GetProductsAdmin() ([]model.Product, error) {

	query := `select p."name", p.description, c."name" as "product_category", b."name" as "product_brand"
	from products p
	inner join categories c on p.category_id  = c.id
	inner join brands b on p.brand_id  = b.id;`

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
			&productObj.Category,
			&productObj.Brand,
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

	query := `insert into Products (name,price,description,category_id,brand_id) values
	($1,$2,$3,$4,$5)`
	
	
	result, err := pr.connection.Exec(query,product.Name,
		product.Price,
		product.Description,
		product.Category,
		product.Brand)	
 	if err != nil {
		panic(err)
	}

	affectedRows, _ := result.RowsAffected()
	fmt.Printf("Rows affeceted: %d", affectedRows)
	return err 
}

func (pr *ProductRepository) GetProductById(id_product int) (*model.Product, error) {

	query := `select p."name", p.description, c."name" as "product_category", b."name" as "product_brand"
	from products p
	inner join categories c on p.category_id  = c.id
	inner join brands b on p.brand_id  = b.id
	where id = $1;`

	result, err := pr.connection.Query(query,id_product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = result.Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
		&produto.Category,
		&produto.Brand,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	result.Close()

	// !! produto com status inativo não é retornado.
	if produto.Product_Status == "inactive"{
		return &produto, nil
	}else{
		return &model.Product{}, nil
	}
}

func (pr *ProductRepository) GetProductByIdAdmin(id_product int) (*model.Product, error) {

	query := `select p."name", p.description, c."name" as "product_category", b."name" as "product_brand"
	from products p
	inner join categories c on p.category_id  = c.id
	inner join brands b on p.brand_id  = b.id
	where id = $1;`

	result, err := pr.connection.Query(query,id_product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = result.Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
		&produto.Category,
		&produto.Brand,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	result.Close()
	return &produto, nil
}

func (pr *ProductRepository) EditProductName(id_product int, value string) (error) {
	var columnName = "name"

	query := 
	`
		update Products 
		set $1 = $2 
		where id = $3
	`

	_, err := pr.connection.Exec(query,columnName,id_product,value)
	if err != nil {
		panic(err)
	}
	return err

}

func (pr *ProductRepository) EditProductPrice(id_product int, value string) (error) {
	var columnName = "price"

	query := 
	`
		update Products 
		set $1 = $2 
		where id = $3
	`

	_, err := pr.connection.Exec(query,columnName,id_product,value)
	if err != nil {
		panic(err)
	}
	return err
	
}

func (pr *ProductRepository) EditProductDescription(id_product int, value string) (error) {
	var columnName = "description"

	query := 
	`
		update Products 
		set $1 = $2 
		where id = $3
	`

	_, err := pr.connection.Exec(query,columnName,id_product,value)
	if err != nil {
		panic(err)
	}
	return err
	
}

func (pr *ProductRepository) RemoveProduct(id_product int) (error) {
	query := `
	Update Products 
	set status = $1
	where id = $2
	`
	
	_, err := pr.connection.Exec(query,id_product); if err != nil {
		panic(err)
	}
	return err 
} 

