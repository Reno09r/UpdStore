package store

type User struct {
	Id       int    `json:"-" db:"user_id"`
	Role     string `json:"role"`
	Fname    string `json:"fname" binding:"required"`
	Lname    string `json:"lname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
