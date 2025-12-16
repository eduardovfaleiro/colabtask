package model

import "time"

type Task struct {
	Base
	ProjectID   ID         `json:"project_id" db:"project_id"`
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	DueDate     *time.Time `json:"due_date" db:"due_date"`
}
