package repository

import (
	"context"
	"database/sql"

	"github.com/eduardovfaleiro/colabtask/internal/model"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
}

type PostgreTaskRepository struct {
	db *sql.DB
}

func NewPostgreTaskRepository(db *sql.DB) *PostgreTaskRepository {
	return &PostgreTaskRepository{db}
}

func (r *PostgreTaskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `INSERT INTO tasks (title, project_id, description, due_date, completed_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.ExecContext(ctx, query, task.Title, task.ProjectID, task.Description, task.DueDate, task.CompletedAt)
	return err
}
