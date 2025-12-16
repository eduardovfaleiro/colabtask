package model

import (
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID

type UserRole string

const (
	RoleAdmin  UserRole = "owner"
	RoleMember UserRole = "member"
)

type Base struct {
	ID        ID        `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
