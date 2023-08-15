package store

type User struct {
	Id       int `json:"-" db:"customer_id"` 
	Fname    string	`json:"rname" binding:"required"`
	Lname    string `json:"lname" binding:"required"`
	Username string	`json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

