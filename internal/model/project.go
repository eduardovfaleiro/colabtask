package model

type Project struct {
	Base
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}
