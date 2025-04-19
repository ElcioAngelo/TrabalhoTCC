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

    query := `
        select p.id, p."name", p.description, p.price, c."name" as "product_category", 
               br."name" as "product_brand", p.status as "product_status"
        from Products p
        inner join product_categories prc on p.id = prc.product_id
        inner join brands br on br.id = p.brand_id
        inner join categories c on c.id = prc.category_id;
    `

    rows, err := pr.connection.Query(query)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return []model.Product{}, err
    }
    defer rows.Close()

    var productList []model.Product
    var productObj model.Product

    for rows.Next() {
        err = rows.Scan(
            &productObj.ID,
            &productObj.Name,
            &productObj.Description,
            &productObj.Price,
            &productObj.Category,
            &productObj.Brand,
            &productObj.Product_Status, // Ensure you're scanning the status
        )
        if err != nil {
            fmt.Println("Error scanning row:", err)
            return []model.Product{}, err
        }

        // Log the values to see what's being returned from the query
        fmt.Printf("Scanned product: %+v\n", productObj)

        // If the status is "inactive", skip adding the product to the list
        if productObj.Product_Status != "inactive" {
            productList = append(productList, productObj)
        } else {
            fmt.Printf("Skipping inactive product: %+v\n", productObj)
        }
    }

    // Check if we encountered any rows
    if len(productList) == 0 {
        fmt.Println("No products found or all products are inactive")
    }

    return productList, nil
}



func (pr *ProductRepository) GetProductsAll() ([]model.Product, error) {

    query := `
        select p.id, p."name", p.description, p.price, c."name" as "product_category", 
               br."name" as "product_brand", p.status as "product_status"
        from Products p
        inner join product_categories prc on p.id = prc.product_id
        inner join brands br on br.id = p.brand_id
        inner join categories c on c.id = prc.category_id;
    `

    rows, err := pr.connection.Query(query)
    if err != nil {
        fmt.Println("Error executing query:", err)
        return []model.Product{}, err
    }
    defer rows.Close()

    var productList []model.Product
    var productObj model.Product

    for rows.Next() {
        err = rows.Scan(
            &productObj.ID,
            &productObj.Name,
            &productObj.Description,
            &productObj.Price,
            &productObj.Category,
            &productObj.Brand,
            &productObj.Product_Status, // Ensure you're scanning the status
        )
        if err != nil {
            fmt.Println("Error scanning row:", err)
            return []model.Product{}, err
        }

        // Log the values to see what's being returned from the query
        fmt.Printf("Scanned product: %+v\n", productObj)

        // If the status is "inactive", skip adding the product to the list
            productList = append(productList, productObj)
    }

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

	query := ` 
 		select p.id, p."name", p.description, p.price, c."name" as "product_category", 
               br."name" as "product_brand", p.status as "product_status"
        from Products p
        inner join product_categories prc on p.id = prc.product_id
        inner join brands br on br.id = p.brand_id
        inner join categories c on c.id = prc.category_id
        where p.id = $1;`

	result, err := pr.connection.Query(query,id_product)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	if result.Next(){
		err = result.Scan(
			&produto.ID,
			&produto.Name,
			&produto.Description,
			&produto.Price,
			&produto.Category,
			&produto.Brand,
			&produto.Product_Status,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}
	}

	result.Close()

	// // !! produto com status inativo não é retornado.
	// if produto.Product_Status == "inactive"{
	// 	return &produto, nil
	// }else{
	// 	return &model.Product{}, nil
	// }
	return &produto, err
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

