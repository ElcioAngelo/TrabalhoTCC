package repository

import (
	"database/sql"
	"fmt"

	"github.comElcioAngelo/TrabalhoTCC.git/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (pr *UserRepository) GetUser() (model.User, error) {

	query := `SELECT id,name,email,cellphone_number,shipping_adress,payment_address from "users" LIMIT 1`
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}

	var user model.User

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CellphoneNumber,
			&user.PaymentAddress,
			&user.ShippingAddress,
		)
			if err != nil {
				fmt.Println(err)
				return model.User{}, err
			}else{
				return model.User{}, fmt.Errorf("no user found")
			}
	}
	rows.Close()
	return user, nil
}



