package model

import (
	"github.com/google/uuid"
)

type ID = uuid.UUID

type UserRole string

const (
	RoleAdmin  UserRole = "owner"
	RoleMember UserRole = "member"
)
