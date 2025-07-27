package adapter

import (
	"context"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/model"
	"github.com/MarrRoss/go-links-to-zip/internal/observability"
	"sync"
)

type TaskRepositoryImpl struct {
	observer *observability.Observability
	tasks    sync.Map
}

func NewTaskRepositoryImpl(
	observer *observability.Observability,
) (*TaskRepositoryImpl, error) {
	return &TaskRepositoryImpl{
		observer: observer,
	}, nil
}

func (repo *TaskRepositoryImpl) CreateTask(
	ctx context.Context,
	task *model.Task,
) error {
	repo.tasks.Store(task.ID, task)
	return nil
}
