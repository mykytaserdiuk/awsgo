package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/aws-go/internal"
	"github.com/mykytaserdiuk/aws-go/pkg/models"
)

type Repository struct {
}

func NewRepo() internal.Repository {
	return &Repository{}
}

func (r *Repository) Create(ctx context.Context, tx sqlx.ExtContext, id, topic, description string, time time.Time) error {
	query := `INSERT INTO todos (id, topic, description, create_time) VALUES ($1,$2, $3, $4)`
	_, err := tx.ExecContext(ctx, query, id, topic, description, time)

	return err
}
func (r *Repository) Delete(ctx context.Context, tx sqlx.ExtContext, id string) error {
	query := `DELETE FROM todos WHERE id = $1`
	_, err := tx.ExecContext(ctx, query, id)

	return err
}
func (r *Repository) Update(ctx context.Context, tx sqlx.ExtContext, id, topic, description string) error {
	query := `UPDATE todos topic = $1, description = $2 WHERE id = $3`
	_, err := tx.ExecContext(ctx, query, topic, description, id)
	return err
}
func (r *Repository) GetByID(ctx context.Context, tx sqlx.ExtContext, id string) (*models.Todo, error) {
	query := `SELECT id, topic, description, create_time FROM todos WHERE id = $1`

	var todo models.Todo
	err := tx.QueryRowxContext(ctx, query, id).StructScan(&todo)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
func (r *Repository) GetAll(ctx context.Context, tx sqlx.ExtContext) ([]*models.Todo, error) {
	todos := make([]*models.Todo, 0)

	query := `SELECT id, topic, description, create_time FROM todos`

	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var todo models.Todo
		rows.StructScan(&todo)
		if todo.ID != "" {
			todos = append(todos, &todo)
		}
	}
	return todos, nil
}
