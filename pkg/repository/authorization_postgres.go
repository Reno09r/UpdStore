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
	query := "SELECT COUNT(*) FROM admins WHERE admin_id = $1"
	var count int
	err := r.db.Get(&count, query, userId)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
