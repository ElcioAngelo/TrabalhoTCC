package repository

import (
	"database/sql"
	"fmt"

	"github.comElcioAngelo/TrabalhoTCC.git/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUser(user_id string) (model.User, error) {

	query := 
	`
	select u.name, u.email, u.cellphone_number, u.shipping_adress, u.payment_address from users u 
	where u.id = $1;
	`

	var user model.User

	err := ur.connection.QueryRow(query, user_id).Scan(
		&user.Name,
		&user.Email,
		&user.CellphoneNumber,
		&user.ShippingAddress,
		&user.PaymentAddress,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, err 
	}
	return user, nil
}

func (ur *UserRepository) CreateUser(user model.User) (error) {
	
	query := `
	Insert into "users" (name,
	email,password,cellphone_number,shipping_adress,payment_address)
	values ($1, $2, $3, $4, $5, $6)
	`

	// * Gera o HASH da senha 
	EncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),
	bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	result,err := ur.connection.Exec(query,user.Name,user.Email, EncryptedPassword,user.CellphoneNumber,
		user.ShippingAddress,user.PaymentAddress)
	if err != nil {
		panic(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf(`Rows affected: %d`,rowsAffected)

	return err
}

// * Estrutura para editar o usuário, necessaria para alterar alguns atributos sem alterar TODOS.

type userEdit struct {
	ID              *int       `json:"id"`                // * Unique identifier for the user
	Name            *string    `json:"name"`              // * Name of the user
	Email           *string    `json:"email"`             // * Email of the user
	Password        *string    `json:"password"`          // * User's password (in a real application, ensure it's hashed)
	CellphoneNumber *string    `json:"cellphone_number"`  // * Cellphone number of the user
	ShippingAddress *string    `json:"shipping_address"`  // * Shipping address for delivery
	PaymentAddress  *string    `json:"payment_address"`
}

func (ur *UserRepository) EditUser(user userEdit) (error){
	//TODO: Terminar essa parte.
	query := `
	Insert into "users" (name,
	email,password,cellphone_number,shipping_adress,payment_address)
	values ($1, $2, $3, $4, $5, $6)
	`

	result,err := ur.connection.Exec(query,user.Name,user.Email, user.Password,user.CellphoneNumber,
		user.ShippingAddress,user.PaymentAddress)
	if err != nil {
		panic(err)
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf(`Rows affected: %d`,rowsAffected)

	return err
}





