package internal

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/aws-go/pkg/models"
)

// Implement Repository inside
type Service interface {
	Create(ctx context.Context, topic, description string) (string, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id, topic, description string) error
	GetByID(ctx context.Context, id string) (*models.Todo, error)
	GetAll(ctx context.Context) ([]*models.Todo, error)
}

//go:generate moq -out repository/repository_moq.go . Repository
type Repository interface {
	Create(ctx context.Context, tx sqlx.ExtContext, id, topic, description string, time time.Time) error
	Delete(ctx context.Context, tx sqlx.ExtContext, id string) error
	Update(ctx context.Context, tx sqlx.ExtContext, id, topic, description string) error
	GetByID(ctx context.Context, tx sqlx.ExtContext, id string) (*models.Todo, error)
	GetAll(ctx context.Context, tx sqlx.ExtContext) ([]*models.Todo, error)
}
