package repository

import (
	"context"
	"database/sql"

	"github.com/eduardovfaleiro/colabtask/internal/model"
)

type ProjectRepository interface {
	Create(ctx context.Context, task *model.Project) error
}

type PostgreProjectRepository struct {
	db *sql.DB
}

func NewPostgreProjectRepository(db *sql.DB) *PostgreProjectRepository {
	return &PostgreProjectRepository{db}
}

func (r *PostgreProjectRepository) Create(ctx context.Context, task *model.Project) error {
	query := `INSERT INTO projects (title, description) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, task.Title, task.Description)
	return err
}
