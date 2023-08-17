package repository

import (
	"github.com/jmoiron/sqlx"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthorizationPostgresSQL(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{db: db}
}

func (r *AuthorizationPostgres) CurrentUserIsAdmin(userId int) (bool, error) {
	query := "SELECT role_name FROM roles WHERE role_id IN (SELECT role_id FROM users WHERE user_id = $1)"
	var roleName string
	err := r.db.Get(&roleName, query, userId)
	if err != nil {
		return false, err
	}

	return roleName == "admin", nil
}


func (r *AuthorizationPostgres) CurrentUserIsCustomer(userId int) (bool, error) {
	query := "SELECT role_name FROM roles WHERE role_id IN (SELECT role_id FROM users WHERE user_id = $1)"
	var roleName string
	err := r.db.Get(&roleName, query, userId)
	if err != nil {
		return false, err
	}

	return roleName == "customer", nil
}