package model

type User struct {
	Base
	Name         string `json:"name" db:"name"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"-" db:"password"`
}
