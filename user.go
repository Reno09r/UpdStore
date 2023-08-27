package store

import "errors"

type User struct {
	Id       int    `json:"-" db:"user_id"`
	Role     string `json:"role" binding:"required" db:"role_id"`
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

func (i UpdateUserInput) Validate() error {
	if i.Fname == nil && i.Lname == nil && i.Username == nil && i.NewPassword == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
