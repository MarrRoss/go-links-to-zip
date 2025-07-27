package adapter

import (
	"context"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/model"
	"github.com/MarrRoss/go-links-to-zip/internal/observability"
	"sync"
)

type TaskRepositoryImpl struct {
	observer *observability.Observability
	tasks    map[model.ID]*model.Task
	mu       sync.RWMutex
}

func NewTaskRepositoryImpl(
	observer *observability.Observability,
) (*TaskRepositoryImpl, error) {
	return &TaskRepositoryImpl{
		observer: observer,
		tasks:    make(map[model.ID]*model.Task),
	}, nil
}

func (repo *TaskRepositoryImpl) CreateTask(
	ctx context.Context,
	task *model.Task,
) error {

}
