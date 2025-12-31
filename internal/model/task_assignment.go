package model

import "time"

type TaskAssignment struct {
	ID         ID        `json:"id" db:"id"`
	TaskID     ID        `json:"task_id" db:"task_id"`
	UserID     ID        `json:"user_id" db:"user_id"`
	AssignedAt time.Time `json:"assigned_at" db:"assigned_at"`
}
