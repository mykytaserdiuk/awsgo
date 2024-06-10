package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/aws-go/internal/repository"
	"github.com/mykytaserdiuk/aws-go/internal/service"
	"github.com/mykytaserdiuk/aws-go/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestUpdateSuccess(t *testing.T) {
	topic := `Old topic`
	description := `Old description`
	time := time.Now()
	uuid := uuid.NewString()
	todo := models.Todo{
		ID:           uuid,
		CreationTime: time,
		Topic:        topic,
		Description:  description,
	}
	newTopic := `New topic`
	newDescription := `New Description`

	mockRepo := repository.RepositoryMock{
		UpdateFunc: func(ctx context.Context, tx sqlx.ExtContext, id, topic, description string) error {
			todo.Description = description
			todo.Topic = topic
			return nil
		},
		GetByIDFunc: func(ctx context.Context, tx sqlx.ExtContext, id string) (*models.Todo, error) {
			if id != uuid {
				return nil, errors.New("Id is not exists")
			}
			return &todo, nil
		},
	}
	service := service.New(nil, &mockRepo)
	ctx := context.Background()
	err := service.Update(ctx, uuid, newTopic, newDescription)

	require.NoError(t, err)
	to, err := service.GetByID(ctx, uuid)
	require.NoError(t, err)
	require.NotNil(t, to)
	require.Equal(t, to, &todo)
}

func TestUpdateFailureWithExistID(t *testing.T) {
	topic := `Old topic`
	description := `Old description`
	time := time.Now()
	uid := uuid.NewString()
	notExistsUuid := uuid.NewString()
	todo := models.Todo{
		ID:           uid,
		CreationTime: time,
		Topic:        topic,
		Description:  description,
	}
	newTopic := `New topic`
	newDescription := `New Description`

	mockRepo := repository.RepositoryMock{
		UpdateFunc: func(ctx context.Context, tx sqlx.ExtContext, id, topic, description string) error {
			todo.Description = description
			todo.Topic = topic
			return nil
		},
		GetByIDFunc: func(ctx context.Context, tx sqlx.ExtContext, id string) (*models.Todo, error) {
			if id != uid {
				return nil, models.ErrTodoNotFound
			}
			return &todo, nil
		},
	}
	service := service.New(nil, &mockRepo)
	ctx := context.Background()
	err := service.Update(ctx, notExistsUuid, newTopic, newDescription)
	require.EqualError(t, err, models.ErrTodoNotFound.Error())
}
