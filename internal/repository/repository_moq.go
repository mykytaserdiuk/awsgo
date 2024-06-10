package repository

import (
	"context"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mykytaserdiuk/aws-go/internal"
	"github.com/mykytaserdiuk/aws-go/pkg/models"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ internal.Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked Repository
//		mockedRepository := &RepositoryMock{
//			CreateFunc: func(ctx context.Context, tx sqlx.ExtContext, id string, topic string, description string, timeMoqParam time.Time) error {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(ctx context.Context, tx sqlx.ExtContext, id string) error {
//				panic("mock out the Delete method")
//			},
//			GetAllFunc: func(ctx context.Context, tx sqlx.ExtContext) ([]*models.Todo, error) {
//				panic("mock out the GetAll method")
//			},
//			GetByIDFunc: func(ctx context.Context, tx sqlx.ExtContext, id string) (*models.Todo, error) {
//				panic("mock out the GetByID method")
//			},
//			UpdateFunc: func(ctx context.Context, tx sqlx.ExtContext, id string, topic string, description string) error {
//				panic("mock out the Update method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, tx sqlx.ExtContext, id string, topic string, description string, timeMoqParam time.Time) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, tx sqlx.ExtContext, id string) error

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context, tx sqlx.ExtContext) ([]*models.Todo, error)

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(ctx context.Context, tx sqlx.ExtContext, id string) (*models.Todo, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, tx sqlx.ExtContext, id string, topic string, description string) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx sqlx.ExtContext
			// ID is the id argument value.
			ID string
			// Topic is the topic argument value.
			Topic string
			// Description is the description argument value.
			Description string
			// TimeMoqParam is the timeMoqParam argument value.
			TimeMoqParam time.Time
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx sqlx.ExtContext
			// ID is the id argument value.
			ID string
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx sqlx.ExtContext
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx sqlx.ExtContext
			// ID is the id argument value.
			ID string
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx sqlx.ExtContext
			// ID is the id argument value.
			ID string
			// Topic is the topic argument value.
			Topic string
			// Description is the description argument value.
			Description string
		}
	}
	lockCreate  sync.RWMutex
	lockDelete  sync.RWMutex
	lockGetAll  sync.RWMutex
	lockGetByID sync.RWMutex
	lockUpdate  sync.RWMutex
}

// Create calls CreateFunc.
func (mock *RepositoryMock) Create(ctx context.Context, tx sqlx.ExtContext, id string, topic string, description string, timeMoqParam time.Time) error {
	if mock.CreateFunc == nil {
		panic("RepositoryMock.CreateFunc: method is nil but Repository.Create was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		Tx           sqlx.ExtContext
		ID           string
		Topic        string
		Description  string
		TimeMoqParam time.Time
	}{
		Ctx:          ctx,
		Tx:           tx,
		ID:           id,
		Topic:        topic,
		Description:  description,
		TimeMoqParam: timeMoqParam,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, tx, id, topic, description, timeMoqParam)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedRepository.CreateCalls())
func (mock *RepositoryMock) CreateCalls() []struct {
	Ctx          context.Context
	Tx           sqlx.ExtContext
	ID           string
	Topic        string
	Description  string
	TimeMoqParam time.Time
} {
	var calls []struct {
		Ctx          context.Context
		Tx           sqlx.ExtContext
		ID           string
		Topic        string
		Description  string
		TimeMoqParam time.Time
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *RepositoryMock) Delete(ctx context.Context, tx sqlx.ExtContext, id string) error {
	if mock.DeleteFunc == nil {
		panic("RepositoryMock.DeleteFunc: method is nil but Repository.Delete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Tx  sqlx.ExtContext
		ID  string
	}{
		Ctx: ctx,
		Tx:  tx,
		ID:  id,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, tx, id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedRepository.DeleteCalls())
func (mock *RepositoryMock) DeleteCalls() []struct {
	Ctx context.Context
	Tx  sqlx.ExtContext
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		Tx  sqlx.ExtContext
		ID  string
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *RepositoryMock) GetAll(ctx context.Context, tx sqlx.ExtContext) ([]*models.Todo, error) {
	if mock.GetAllFunc == nil {
		panic("RepositoryMock.GetAllFunc: method is nil but Repository.GetAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Tx  sqlx.ExtContext
	}{
		Ctx: ctx,
		Tx:  tx,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc(ctx, tx)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//
//	len(mockedRepository.GetAllCalls())
func (mock *RepositoryMock) GetAllCalls() []struct {
	Ctx context.Context
	Tx  sqlx.ExtContext
} {
	var calls []struct {
		Ctx context.Context
		Tx  sqlx.ExtContext
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *RepositoryMock) GetByID(ctx context.Context, tx sqlx.ExtContext, id string) (*models.Todo, error) {
	if mock.GetByIDFunc == nil {
		panic("RepositoryMock.GetByIDFunc: method is nil but Repository.GetByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Tx  sqlx.ExtContext
		ID  string
	}{
		Ctx: ctx,
		Tx:  tx,
		ID:  id,
	}
	mock.lockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	mock.lockGetByID.Unlock()
	return mock.GetByIDFunc(ctx, tx, id)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//
//	len(mockedRepository.GetByIDCalls())
func (mock *RepositoryMock) GetByIDCalls() []struct {
	Ctx context.Context
	Tx  sqlx.ExtContext
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		Tx  sqlx.ExtContext
		ID  string
	}
	mock.lockGetByID.RLock()
	calls = mock.calls.GetByID
	mock.lockGetByID.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *RepositoryMock) Update(ctx context.Context, tx sqlx.ExtContext, id string, topic string, description string) error {
	if mock.UpdateFunc == nil {
		panic("RepositoryMock.UpdateFunc: method is nil but Repository.Update was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		Tx          sqlx.ExtContext
		ID          string
		Topic       string
		Description string
	}{
		Ctx:         ctx,
		Tx:          tx,
		ID:          id,
		Topic:       topic,
		Description: description,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, tx, id, topic, description)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//
//	len(mockedRepository.UpdateCalls())
func (mock *RepositoryMock) UpdateCalls() []struct {
	Ctx         context.Context
	Tx          sqlx.ExtContext
	ID          string
	Topic       string
	Description string
} {
	var calls []struct {
		Ctx         context.Context
		Tx          sqlx.ExtContext
		ID          string
		Topic       string
		Description string
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
