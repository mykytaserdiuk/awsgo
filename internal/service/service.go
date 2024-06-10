package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/aws-go/internal"
	"github.com/mykytaserdiuk/aws-go/pkg/models"
)

// Service without bussines logic
type Service struct {
	db   *sqlx.DB
	repo internal.Repository
}

func New(db *sqlx.DB, repo internal.Repository) internal.Service {
	return &Service{db, repo}
}

func (s *Service) Create(ctx context.Context, topic, description string) (string, error) {
	id := uuid.NewString()
	time := time.Now().UTC()
	err := s.repo.Create(ctx, s.db, id, topic, description, time)
	if err != nil {
		return "", err
	}
	return id, nil
}
func (s *Service) Delete(ctx context.Context, id string) error {
	todo, err := s.repo.GetByID(ctx, s.db, id)
	if err != nil {
		return err
	}
	if todo == nil {
		return models.ErrTodoNotFound
	}

	err = s.repo.Delete(ctx, s.db, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) Update(ctx context.Context, id, topic, description string) error {
	todo, err := s.repo.GetByID(ctx, s.db, id)
	if err != nil {
		return err
	}
	if todo == nil {
		return models.ErrTodoNotFound
	}

	err = s.repo.Update(ctx, s.db, id, topic, description)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) GetByID(ctx context.Context, id string) (*models.Todo, error) {
	todo, err := s.repo.GetByID(ctx, s.db, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (s *Service) GetAll(ctx context.Context) ([]*models.Todo, error) {
	todos, err := s.repo.GetAll(ctx, s.db)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
