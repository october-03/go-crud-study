package models

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	password string `json:"-"`
}

func NewUser(username string, email string, password string) User {
	return User{username, email, password}
}
