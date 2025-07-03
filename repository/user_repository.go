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

func (ur *UserRepository) EncryptValue(value string) (string, error) {
	EncryptedValue, err := bcrypt.GenerateFromPassword([]byte(value),
		bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	return string(EncryptedValue), err
}

func (ur *UserRepository) GetUserByID(user_id int) (model.User, error) {

	query :=
		`
	select u.id, u.name, u.email, u.cellphone_number,
	uid.state, uid.postal_code, uid.city, uid.address, uid.address_number,
	ui.user_role 
	from users u
	inner join user_address_information uid on uid.user_id = u.id
	inner join user_information ui on ui.user_id = u.id
	where u.id = $1;
	`

	var user model.User

	err := ur.connection.QueryRow(query, user_id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CellphoneNumber,
		&user.State,
		&user.City,
		&user.PostalCode,
		&user.Address,
		&user.AddressNumber,
		&user.UserRole,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Error: %v", err)
		}
		return user, err
	}
	return user, nil
}

// ? Estrutura para evitar de retornar a senha do usuário.
type ReturnUser struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	CellphoneNumber string `json:"cellphone_number"`
	State           string `json:"state"`
	PostalCode      string `json:"postal_code"`
	City            string `json:"city"`
	Address         string `json:"address"`
	AddressNumber   int    `json:"address_number"`
	UserRole        string `json:"user_role"`
	UserStatus      string `json:"user_status"`
}

func (ur *UserRepository) GetUsers() ([]ReturnUser, error) {

	query := `select u.id,
			   u.name,
			   u.email,
			   u.cellphone_number,
			   uai.state,
			   uai.postal_code,
			   uai.city,
			   uai.address,
			   uai.address_number,
			   ui.user_role,
			   ui.status
		from users u
		join user_information ui on ui.user_id = u.id
		inner join user_address_information uai on uai.user_id = u.id;
		`

	rows, err := ur.connection.Query(query)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	defer rows.Close()

	var userList []ReturnUser

	for rows.Next() {
		var UserObj ReturnUser

		err = rows.Scan(
			&UserObj.ID,
			&UserObj.Name,
			&UserObj.Email,
			&UserObj.CellphoneNumber,
			&UserObj.State,
			&UserObj.PostalCode,
			&UserObj.City,
			&UserObj.Address,
			&UserObj.AddressNumber,
			&UserObj.UserRole,
			&UserObj.UserStatus,
		)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		userList = append(userList, UserObj)
		// Log the values to see what's being returned from the query
		fmt.Printf("Scanned User: %+v\n", UserObj)
	}
	// Check if we encountered any rows
	if len(userList) == 0 {
		fmt.Println("No users found")
	}

	return userList, nil
}

type User struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	CellphoneNumber string `json:"cellphone_number"`
	State           string `json:"state"`
	PostalCode      string `json:"postal_code"`
	City            string `json:"city"`
	Address         string `json:"address"`
	AddressNumber   string `json:"address_number"`
	UserRole        string `json:"user_role"`
}

func (ur *UserRepository) UserVerification(user_email string, user_password string) (User, error) {

	query := `
		select u.id,
			   u.name,
			   u.email,
			   u.cellphone_number,
			   uai.state,
			   uai.postal_code,
			   uai.city,
			   uai.address,
			   uai.address_number,
			   ui.user_role,
			   u.password -- Fetch the password hash from the database
		from users u
		join user_information ui on ui.user_id = u.id
		inner join user_address_information uai on uai.user_id = u.id
		where u.email = $1;
		`

	var user User

	err := ur.connection.QueryRow(query, user_email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CellphoneNumber,
		&user.State,
		&user.PostalCode,
		&user.City,
		&user.Address,
		&user.AddressNumber,
		&user.UserRole,
		&user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle user not found
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user_password))
	if err != nil {
		return user, fmt.Errorf("invalid password")
	}

	return user, nil
}

func (ur *UserRepository) CreateUser(user model.User) error {
	// * Usando uma Common Table Expression, é possivel realizar várias
	// * inserções sem necessitar de várias váriaveis.
	query :=
		`
	WITH inserted_user AS (
	INSERT INTO users (name, email, password, cellphone_number)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	),
	inserted_address AS (
	INSERT INTO user_address_information (
		user_id, state, postal_code, city, address, address_number
	)
	VALUES (
		(SELECT id FROM inserted_user), $5, $6, $7, $8, $9
	)
	),
	inserted_user_info AS (
	INSERT INTO user_information (
		user_id, creation_date,update_date,user_role,status)
	VALUES (
		(SELECT id FROM inserted_user), now(), NULL,'customer','active')
	)
	SELECT id FROM inserted_user;
	`

	// ? Função para gerar o hash criptografado da senha
	EncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),
		bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	_, err = ur.connection.Exec(query, user.Name, user.Email, EncryptedPassword, user.CellphoneNumber,
		user.State, user.PostalCode, user.City, user.Address, user.AddressNumber)
	if err != nil {
		return fmt.Errorf("error while creating user: %w", err)
	}
	return err
}

func (ur *UserRepository) UpdateUser(user model.User) error {
	query :=
		`
		WITH updated_user AS (
		UPDATE users
		SET name = $1,
			email = $2,
			password = $3,
			cellphone_number = $4
		WHERE id = $5
		returning id
		),
		updated_address AS (
		UPDATE user_address_information
		SET state = $6,
			postal_code = $7,
			city = $8,
			address = $9,
			address_number = $10
		WHERE user_id = (SELECT id FROM updated_user)
		returning *
		)
		select 'update succesfull!';
	`

	// // ? Função para gerar o hash criptografado da senha
	// EncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),
	// 	bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err.Error())
	// }

	_, err := ur.connection.Exec(
		query,
		user.Name,
		user.Email,
		user.CellphoneNumber,
		user.ID,
		user.State,
		user.PostalCode,
		user.City,
		user.Address,
		user.AddressNumber,
	)
	if err != nil {
		return fmt.Errorf("error while creating user: %w", err)
	}

	return err
}

func (ur *UserRepository) setUserProducts(user_id int, order_id int, product_id int, quantity int) error {

	query :=
		`
	insert into user_order (user_id, quantity, product_id, order_id)
	values ($1, $2, $3, $4);
	`

	_, err := ur.connection.Exec(query, user_id, quantity, product_id, order_id)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	return err
}

func (ur *UserRepository) RemoveUser(user_id int) error {
	query :=
		`
		update users
		set user_status = 'disabled'
		where id = $1
	`

	_, err := ur.connection.Exec(query, user_id)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	return err
}
