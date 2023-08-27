package repository

import (
	"errors"
	"fmt"
	"strings"

	store "github.com/Reno09r/Store"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Get(userId int) (store.User, error) {
	var user store.User
	query := fmt.Sprintf("SELECT user_fname, user_lname, username FROM %s WHERE user_id = $1", UsersTable)
	err := r.db.Get(&user, query, userId)
	return user, err
}

func (r *UserPostgres) Delete(userId int) error {
	var user store.User
	queryCheck := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", UsersTable)
	err := r.db.Get(&user, queryCheck, userId)
	if err != nil {
		return errors.New("Delete by non-existent userId")
	}
	var ids []int
	query := "SELECT purchase_id FROM purchases WHERE user_id = $1"
	err = r.db.Select(&ids, query, userId)
	if err != nil {
		return err
	}
	for _, id := range ids {
		query := "DELETE FROM purchase_items WHERE purchase_id = $1"
		_, err = r.db.Exec(query, id)
		if err != nil {
			return err
		}
	}
	query = "DELETE FROM purchases WHERE user_id = $1"
	_, err = r.db.Exec(query, userId)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", UsersTable)
	_, err = r.db.Exec(query, userId)
	return err
}

func (r *UserPostgres) Update(userId int, input store.UpdateUserInput) error {
	var user store.User
	userCheck := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", UsersTable)
	err := r.db.Get(&user, userCheck, userId)
	if err != nil {
		return errors.New("Update by non-existent userId")
	}
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Fname != nil {
		setValues = append(setValues, fmt.Sprintf("user_fname=$%d", argId))
		args = append(args, *input.Fname)
		argId++
	}

	if input.Lname != nil {
		setValues = append(setValues, fmt.Sprintf("user_lname=$%d", argId))
		args = append(args, *input.Lname)
		argId++
	}

	if input.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *input.Username)
		argId++
	}

	if input.NewPassword != nil {
		var hashed_password string
		passwordCheck := fmt.Sprintf("SELECT hashed_password FROM %s WHERE user_id = $1", UsersTable)
		err := r.db.Get(&hashed_password, passwordCheck, userId)
		if err != nil {
			return err
		}
		if input.OldPassword == nil {
			return errors.New("old password was not written")
		}
		if bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(*input.OldPassword)) != nil {
			return errors.New("incorrect old password")
		}
		if bcrypt.CompareHashAndPassword([]byte(*input.NewPassword), []byte(*input.OldPassword)) == nil{
			return errors.New("old and new passwords are the same")
		}
		setValues = append(setValues, fmt.Sprintf("hashed_password=$%d", argId))
		args = append(args, *input.NewPassword)
		argId++
	} else if input.NewPassword == nil && input.OldPassword != nil{
		return errors.New("new password was not written")
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ct SET %s WHERE ct.user_id =$%d",
		UsersTable, setQuery, argId)

	_, err = r.db.Exec(query, append(args, userId)...)
	return err
}
