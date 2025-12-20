package repository

import (
	"context"
	"database/sql"

	"github.com/eduardovfaleiro/colabtask/internal/model"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	GetByID(ctx context.Context, id model.ID) (*model.Task, error)
}

type PostgreTaskRepository struct {
	db *sql.DB
}

func NewPostgreTaskRepository(db *sql.DB) *PostgreTaskRepository {
	return &PostgreTaskRepository{db}
}

func (r *PostgreTaskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `INSERT INTO tasks (title, description, due_date) VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, task.ID, task.Title, task.Description, task.DueDate)
	return err
}
