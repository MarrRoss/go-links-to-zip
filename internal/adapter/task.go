package adapter

import (
	"context"
	"fmt"
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

func (repo *TaskRepositoryImpl) CheckActiveTasksLimit(
	ctx context.Context,
) error {
	var activeTasks int
	repo.tasks.Range(func(key, value any) bool {
		task, ok := value.(*model.Task)
		if !ok {
			return true
		}
		if task.Status != model.StatusFinished && task.Status != model.StatusError {
			activeTasks++
			if activeTasks >= 3 {
				return false
			}
		}
		return true
	})

	if activeTasks >= 3 {
		return fmt.Errorf("%w: %w", ErrTaskLimit, ErrStorage)
	}
	return nil
}

func (repo *TaskRepositoryImpl) PrintAllTasks() {
	repo.tasks.Range(func(key, value any) bool {
		task, ok := value.(*model.Task)
		if !ok {
			fmt.Println("invalid type stored in tasks map")
			return true
		}
		fmt.Printf("Task ID: %s, Status: %s\n", task.ID.String(), task.Status)
		return true
	})
}
