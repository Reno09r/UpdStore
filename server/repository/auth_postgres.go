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

func (r *AuthpostgresSQL) CreateUser(user store.User) (int, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (customer_fname, customer_lname, username, hashed_password) VALUES ($1, $2, $3, $4) RETURNING customer_id", UsersTable)
	row := r.db.QueryRow(query, user.Fname, user.Lname, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *AuthpostgresSQL) GetUser(username, password string) (store.User, error) {
	var user store.User
	query := fmt.Sprintf("SELECT hashed_password FROM %s WHERE username=$1", UsersTable)
	row := r.db.QueryRow(query, username)
	var hashed_password string
	if err := row.Scan(&hashed_password); err != nil {
		return user, err
	}
	if bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password)) != nil {
		return user, errors.New("Неверный пароль")
	}
	query = fmt.Sprintf("SELECT customer_id FROM %s WHERE username=$1 AND hashed_password=$2", UsersTable)
	err := r.db.Get(&user, query, username, hashed_password)
	return user, err
}

