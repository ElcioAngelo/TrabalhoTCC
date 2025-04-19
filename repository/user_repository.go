package repository

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"trabalhoTcc.com/mod/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) EncryptValue(value string) (string) {
	EncryptedValue, err := bcrypt.GenerateFromPassword([]byte(value),
	bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(EncryptedValue)
}

func (ur *UserRepository) GetUserByID(user_id string) (model.User, error) {

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

type User struct {
	ID              int       `json:"id"`                
	Name            string    `json:"name"`              
	Email           string    `json:"email"`            
	CellphoneNumber string    `json:"cellphone_number"`  
	ShippingAddress string    `json:"shipping_address"`  
	PaymentAddress  string    `json:"payment_address"`      
	UserRole		string 	  `json:"user_role"`
}
func (ur *UserRepository) UserVerification(user_email string, user_password string) (User, error) {

	query := 
	`
	select u.id,
	u.name,
	u.email,
    u.cellphone_number,	    
    u.shipping_adress,
	u.payment_address,
	ui.user_role 
	from users u 
	join user_information ui on ui.user_id = u.id
	where u.email = $1 and u.password = $2;
	
	`
	
	var user User

	// encryptedPassword := ur.EncryptValue(user_password)

	err := ur.connection.QueryRow(query, user_email, user_password).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CellphoneNumber,
		&user.PaymentAddress,
		&user.ShippingAddress,
		&user.UserRole,
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
	
	query := 
	`
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


func(ur *UserRepository) EditUserName(id int, value string) (error){
		
	_, err := ur.connection.Exec(`update Users set name = $1 where id = $2`,value,id)
	if err != nil {
		panic(err)
	}
	return err
}


func(ur *UserRepository) EditUserPassword(id int, value string) (error){
	
	
	query := 
	`
		update Users 
		set password = $1
		where id = $2
	`

	result,err := ur.connection.Exec(query,ur.EncryptValue(value),id)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	return err
}


func(ur *UserRepository) EditUserCellphone(id int, value string) (error){	
	query := 
	`
		update Users 
		set cellphone_number = $1
		where id = $2
	`

	_, err := ur.connection.Exec(query,value,id)
	if err != nil {
		panic(err)
	}
	return err
}


func(ur *UserRepository) EditUserEmail(id int, value string) (error){
		
	query := 
	`
		update Users 
		set email = $1
		where id = $2
	`

	_, err := ur.connection.Exec(query,value,id)
	if err != nil {
		panic(err)
	}
	return err
}

func(ur *UserRepository) EditUserPaymentAdress(id int, value string) (error){
	
	query := 
	`
		update Users 
		set payment_address = $1
		where id = $2
	`

	_, err := ur.connection.Exec(query,value,id)
	if err != nil {
		panic(err)
	}
	return err
}

func(ur *UserRepository) EditUserShipmentAdress(id int, value string) (error){
	
	var columnName = "shipping_adress"
	
	query := 
	`
		update Users 
		set $1 = $2
		where id = $3
	`

	_, err := ur.connection.Exec(query,columnName,value,id)
	if err != nil {
		panic(err)
	}
	return err
}

func(ur *UserRepository) setUserOrder(user_id int) (error){
	
	query := 
	`
		insert into orders (order_date, payment_method, status, user_id)
		values(now(), null, 'pending', $1);
	`

	_, err := ur.connection.Exec(query, user_id);
	if err != nil {
		panic(err.Error())
	}
	return err
}

func(ur *UserRepository) setUserProducts(user_id int, order_id int,product_id int, quantity int) (error) {

	query :=
	`
	insert into user_order (user_id, quantity, product_id, order_id)
	values ($1, $2, $3, $4);
	`

	_, err := ur.connection.Exec(query, user_id, quantity, product_id, order_id);
	if err != nil {
		panic(err.Error())
	}

	return err 
}
