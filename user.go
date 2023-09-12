package store

import "errors"

type Role struct {
	Id   int    `json:"-" db:"role_id"`
	Name string `json:"role" binding:"required" db:"role_name"`
}

type User struct {
	Id int `json:"-" db:"user_id"`
	Role
	Fname    string `json:"fname" binding:"required" db:"user_fname"`
	Lname    string `json:"lname" binding:"required" db:"user_lname"`
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"hashed_password"`
}

type UpdateUserInput struct {
	Fname       *string `json:"fname"`
	Lname       *string `json:"lname"`
	Username    *string `json:"username"`
	OldPassword *string `json:"old_password"`
	NewPassword *string `json:"new_password"`
}

type UserCardInput struct {
	CardNumber       *string `json:"cardnumber" binding:"required"`
	CardExpirationDate       *string `json:"expdate" binding:"required"`
	CVV    *string `json:"cvv" binding:"required"`
}

func (i UpdateUserInput) Validate() error {
	if i.Fname == nil && i.Lname == nil && i.Username == nil && i.NewPassword == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UserCardInput) Validate() error {
	if i.CardNumber == nil && i.CardExpirationDate == nil && i.CVV == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
