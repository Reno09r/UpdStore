package repository

import (
	"errors"
	"fmt"

	"github.com/Reno09r/Store"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AuthpostgresSQL struct {
	db *sqlx.DB
}

func NewAuthPostgresSQL(db *sqlx.DB) *AuthpostgresSQL {
	return &AuthpostgresSQL{db: db}
}

func (r *AuthpostgresSQL) CreateCustomer(user store.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var id int64
	insertUsernameQuery := fmt.Sprintf(`
		INSERT INTO %s (username) 
		VALUES ($1) 
		RETURNING username_id`, UsernamesTable)
	row := tx.QueryRow(insertUsernameQuery, user.Username)
	if err := row.Scan(&user.Id); err != nil {
		return 0, err
	}

	insertUserQuery := fmt.Sprintf(`
		INSERT INTO %s (customer_fname, customer_lname, username_id, hashed_password) 
		VALUES ($1, $2, $3, $4) 
		RETURNING customer_id`, CustomersTable)
	row = tx.QueryRow(insertUserQuery, user.Fname, user.Lname, user.Id, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthpostgresSQL) GetCustomer(username, password string) (store.User, error) {
	var user store.User
	query := fmt.Sprintf(`SELECT hashed_password FROM %s WHERE username_id = (SELECT username_id FROM %s WHERE username = $1);`, CustomersTable, UsernamesTable)
	row := r.db.QueryRow(query, username)
	var hashed_password string
	if err := row.Scan(&hashed_password); err != nil {
		return user, err
	}
	if bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password)) != nil {
		return user, errors.New("incorrect password")
	}
	query = fmt.Sprintf("SELECT customer_id FROM %s WHERE username_id=(SELECT username_id FROM %s WHERE username = $1) AND hashed_password=$2", CustomersTable, UsernamesTable)
	err := r.db.Get(&user, query, username, hashed_password)
	return user, err
}

func (r *AuthpostgresSQL) CreateAdmin(user store.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var id int64
	insertUsernameQuery := fmt.Sprintf(`
		INSERT INTO %s (username) 
		VALUES ($1) 
		RETURNING username_id`, UsernamesTable)
	row := tx.QueryRow(insertUsernameQuery, user.Username)
	if err := row.Scan(&user.Id); err != nil {
		return 0, err
	}

	insertUserQuery := fmt.Sprintf(`
		INSERT INTO %s (admin_fname, admin_lname, username_id, hashed_password) 
		VALUES ($1, $2, $3, $4) 
		RETURNING admin_id`, AdminsTable)
	row = tx.QueryRow(insertUserQuery, user.Fname, user.Lname, user.Id, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *AuthpostgresSQL) GetAdmin(username, password string) (store.User, error) {
	var user store.User
	query := fmt.Sprintf(`SELECT hashed_password FROM %s WHERE username_id = (SELECT username_id FROM %s WHERE username = $1);`, AdminsTable, UsernamesTable)
	row := r.db.QueryRow(query, username)
	var hashed_password string
	if err := row.Scan(&hashed_password); err != nil {
		return user, err
	}
	if bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password)) != nil {
		return user, errors.New("incorrect password")
	}
	query = fmt.Sprintf("SELECT admin_id FROM %s WHERE username_id=(SELECT username_id FROM %s WHERE username = $1) AND hashed_password=$2", AdminsTable, UsernamesTable)
	err := r.db.Get(&user, query, username, hashed_password)
	return user, err
}
