package model

type ProjectCollaborator struct {
	ID        ID       `json:"id" db:"id"`
	UserID    ID       `json:"user_id" db:"user_id"`
	ProjectID ID       `json:"project_id" db:"project_id"`
	Role      UserRole `json:"role" db:"role"`
}
