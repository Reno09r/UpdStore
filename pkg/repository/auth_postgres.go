package repository

import (
	"errors"
	"fmt"

	store "github.com/Reno09r/Store"
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
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	insertUserQuery := fmt.Sprintf(`
		SELECT role_id FROM %s WHERE role_name = $1`, RolesTable)
	row := tx.QueryRow(insertUserQuery, user.Role.Name)
	if err := row.Scan(&user.Role.Id); err != nil {
		return 0, err
	}
	var id int
	insertUserQuery = fmt.Sprintf(`
		INSERT INTO %s (user_fname, user_lname, username, hashed_password, role_id) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING user_id`, UsersTable)
	row = tx.QueryRow(insertUserQuery, user.Fname, user.Lname, user.Username, user.Password, user.Role.Id)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *AuthpostgresSQL) GetUser(username, password string) (store.User, error) {
	var user store.User
	query := fmt.Sprintf(`SELECT hashed_password FROM %s WHERE username = $1;`, UsersTable)
	row := r.db.QueryRow(query, username)
	var hashed_password string
	if err := row.Scan(&hashed_password); err != nil {
		return user, errors.New("user has been not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password)) != nil {
		return user, errors.New("incorrect password")
	}
	query = fmt.Sprintf("SELECT user_id FROM %s WHERE username = $1 AND hashed_password=$2", UsersTable)
	err := r.db.Get(&user, query, username, hashed_password)
	return user, err
}
